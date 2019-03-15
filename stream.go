package dht

import (
	"context"
	"sync"

	pbio "github.com/gogo/protobuf/io"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	inet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	"golang.org/x/xerrors"
)

func (dht *IpfsDHT) newStream(ctx context.Context, p peer.ID) (*stream, error) {
	s, err := dht.host.NewStream(ctx, p, dht.protocols...)
	if err != nil {
		return nil, xerrors.Errorf("opening stream: %w", err)
	}
	ps := &stream{
		stream: s,
		w:      newBufferedDelimitedWriter(s),
		r:      pbio.NewDelimitedReader(s, inet.MessageSizeMax),
		m:      make(chan chan *pb.Message, 1),
	}
	go func() {
		ps.reader()
		dht.streamPool.delete(ps, p)
		ps.reset()
	}()
	return ps, nil
}

type stream struct {
	stream interface {
		Reset() error
	}
	w bufferedWriteCloser
	r pbio.ReadCloser

	// Synchronizes m and readerErr.
	mu sync.Mutex
	// Receives channels to send responses on.
	m         chan chan *pb.Message
	readerErr error
}

func (me *stream) reset() {
	me.stream.Reset()
}

func (me *stream) send(m *pb.Message) (err error) {
	if err := me.w.WriteMsg(m); err != nil {
		return xerrors.Errorf("writing message: %w", err)
	}
	if err := me.w.Flush(); err != nil {
		return xerrors.Errorf("flushing: %w", err)
	}
	return nil
}

func (me *stream) request(ctx context.Context, req *pb.Message) (<-chan *pb.Message, error) {
	replyChan := make(chan *pb.Message, 1)
	me.mu.Lock()
	if err := me.errLocked(); err != nil {
		me.mu.Unlock()
		return nil, err
	}
	select {
	case me.m <- replyChan:
	default:
		me.mu.Unlock()
		panic("message pipeline full")
	}
	me.mu.Unlock()
	err := me.send(req)
	return replyChan, err
}

// Handles the error returned from the read loop.
func (me *stream) reader() {
	err := me.readLoop()
	me.mu.Lock()
	me.readerErr = err
	close(me.m)
	me.mu.Unlock()
	for mc := range me.m {
		close(mc)
	}
}

// Reads from the stream until something is wrong.
func (me *stream) readLoop() error {
	for {
		var m pb.Message
		err := me.r.ReadMsg(&m)
		if err != nil {
			return err
		}
		select {
		case mc := <-me.m:
			mc <- &m
		default:
			return xerrors.New("read superfluous message")
		}
	}
}

func (me *stream) err() error {
	me.mu.Lock()
	defer me.mu.Unlock()
	return me.errLocked()
}

// A stream has gone bad when the reader has given up.
func (me *stream) errLocked() error {
	if me.readerErr != nil {
		return xerrors.Errorf("reader: %w", me.readerErr)
	}
	return nil
}
