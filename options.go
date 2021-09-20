package wire

import (
	"crypto/tls"

	"go.uber.org/zap"
)

// OptionFn options pattern used to define and set options for the given
// PostgreSQL server.
type OptionFn func(*Server)

// SimpleQuery sets the simple query handle inside the given server instance.
func SimpleQuery(fn SimpleQueryFn) OptionFn {
	return func(srv *Server) {
		srv.SimpleQuery = fn
	}
}

// MessageBufferSize sets the message buffer size which is allocated once a new
// connection gets constructed. If a negative value or zero value is provided is
// the default message buffer size used.
func MessageBufferSize(size int) OptionFn {
	return func(srv *Server) {
		srv.BufferedMsgSize = size
	}
}

// Certificates sets the given TLS certificates to be used to initialize a
// secure connection between the front-end (client) and back-end (server).
func Certificates(certs []tls.Certificate) OptionFn {
	return func(srv *Server) {
		srv.Certificates = certs
	}
}

// GlobalParameters sets the server parameters which are send back to the
// front-end (client) once a handshake has been established.
func GlobalParameters(params Parameters) OptionFn {
	return func(srv *Server) {
		srv.Parameters = params
	}
}

// Logger sets the given zap logger as the default logger for the given server.
func Logger(logger *zap.Logger) OptionFn {
	return func(srv *Server) {
		srv.logger = logger
	}
}