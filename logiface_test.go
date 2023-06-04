package bench

import (
	"github.com/joeycumines/ilogrus"
	"github.com/joeycumines/izerolog"
	"github.com/joeycumines/logiface"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func benchLogiface[E logiface.Event](b *testing.B, config func(stream *blackholeStream) logiface.Option[E]) {
	newLogger := func(stream *blackholeStream, options ...logiface.Option[E]) *logiface.Logger[E] {
		return logiface.New(append([]logiface.Option[E]{config(stream)}, options...)...)
	}

	b.Run(`TextPositive`, func(b *testing.B) {
		stream := &blackholeStream{}
		logger := newLogger(stream)
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info().Log("The quick brown fox jumps over the lazy dog")
			}
		})

		if stream.WriteCount() != uint64(b.N) {
			b.Fatalf("Log write count")
		}
	})

	b.Run(`TextNegative`, func(b *testing.B) {
		stream := &blackholeStream{}
		logger := newLogger(stream, logiface.WithLevel[E](logiface.LevelError))
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info().Log("The quick brown fox jumps over the lazy dog")
			}
		})

		if stream.WriteCount() != uint64(0) {
			b.Fatalf("Log write count")
		}
	})

	b.Run(`JSONNegative`, func(b *testing.B) {
		stream := &blackholeStream{}
		logger := newLogger(stream, logiface.WithLevel[E](logiface.LevelError))
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info().
					Str("rate", "15").
					Int("low", 16).
					Float32("high", 123.2).
					Log("The quick brown fox jumps over the lazy dog")
			}
		})

		if stream.WriteCount() != uint64(0) {
			b.Fatalf("Log write count")
		}
	})

	b.Run(`JSONPositive`, func(b *testing.B) {
		stream := &blackholeStream{}
		logger := newLogger(stream)
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info().
					Str("rate", "15").
					Int("low", 16).
					Float32("high", 123.2).
					Log("The quick brown fox jumps over the lazy dog")
			}
		})

		if stream.WriteCount() != uint64(b.N) {
			b.Fatalf("Log write count")
		}
	})

	b.Run(`JSONPositiveInterface`, func(b *testing.B) {
		stream := &blackholeStream{}
		// the only difference from the JSONPositive test case is the .Logger bit
		logger := newLogger(stream).Logger()
		b.ResetTimer()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logger.Info().
					Str("rate", "15").
					Int("low", 16).
					Float32("high", 123.2).
					Log("The quick brown fox jumps over the lazy dog")
			}
		})

		if stream.WriteCount() != uint64(b.N) {
			b.Fatalf("Log write count")
		}
	})
}

func BenchmarkLogifaceZerologJSON(b *testing.B) {
	benchLogiface(b, func(stream *blackholeStream) logiface.Option[*izerolog.Event] {
		return izerolog.L.WithZerolog(zerolog.New(stream).
			With().
			Timestamp().
			Logger())
	})
}

func BenchmarkLogifaceZerologText(b *testing.B) {
	benchLogiface(b, func(stream *blackholeStream) logiface.Option[*izerolog.Event] {
		return izerolog.L.WithZerolog(zerolog.New(stream).
			Output(zerolog.ConsoleWriter{
				Out:     stream,
				NoColor: true,
				// same as logrus' default
				TimeFormat: time.RFC3339,
			}).
			With().
			Timestamp().
			Logger())
	})
}

func BenchmarkLogifaceLogrusJSON(b *testing.B) {
	benchLogiface(b, func(stream *blackholeStream) logiface.Option[*ilogrus.Event] {
		logger := logrus.New()
		logger.Formatter = &logrus.JSONFormatter{}
		logger.Out = stream
		return ilogrus.L.WithLogrus(logger)
	})
}

func BenchmarkLogifaceLogrusText(b *testing.B) {
	benchLogiface(b, func(stream *blackholeStream) logiface.Option[*ilogrus.Event] {
		logger := logrus.New()
		logger.Formatter = &logrus.TextFormatter{
			DisableColors:  true,
			FullTimestamp:  true,
			DisableSorting: true,
		}
		logger.Out = stream
		return ilogrus.L.WithLogrus(logger)
	})
}
