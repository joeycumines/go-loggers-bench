package bench

import (
	izerolog "github.com/joeycumines/go-utilpkg/logiface/zerolog"
	"github.com/rs/zerolog"
	"testing"
)

func BenchmarkLogifaceZerologTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := izerolog.L.New(izerolog.L.WithZerolog(zerolog.New(stream).With().Timestamp().Logger()))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//logger.Info().Msg("The quick brown fox jumps over the lazy dog")
			logger.Info().Log("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLogifaceZerologTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := izerolog.L.New(
		izerolog.L.WithZerolog(zerolog.New(stream).With().Timestamp().Logger()),
		izerolog.L.WithLevel(izerolog.L.LevelError()),
	)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//logger.Info().Msg("The quick brown fox jumps over the lazy dog")
			logger.Info().Log("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLogifaceZerologJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := izerolog.L.New(
		izerolog.L.WithZerolog(zerolog.New(stream).With().Timestamp().Logger()),
		izerolog.L.WithLevel(izerolog.L.LevelError()),
	)
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
}

func BenchmarkLogifaceZerologJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := izerolog.L.New(izerolog.L.WithZerolog(zerolog.New(stream).With().Timestamp().Logger()))
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
}
