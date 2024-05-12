package packets_2023_test

import (
	"fmt"
	"testing"

	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
	f1_2023 "github.com/DaanV2/go-f1-library/packets/2023"
	test_util "github.com/DaanV2/go-f1-library/test/packets/util"
	"github.com/stretchr/testify/require"
)

func Test_F1_2023_Packets(t *testing.T) {
	results := test_util.NewResults()
	err := results.Load(".")
	require.NoError(t, err)

	parser := f1_2023.NewPacketParser()

	t.Run("Packet CarDamage", Parse_Tests(results.PacketCarDamage, parser.ParsePacket))
	t.Run("Packet CarSetups", Parse_Tests(results.PacketCarSetups, parser.ParsePacket))
	t.Run("Packet CarStatus", Parse_Tests(results.PacketCarStatus, parser.ParsePacket))
	t.Run("Packet CarTelemetry", Parse_Tests(results.PacketCarTelemetry, parser.ParsePacket))
	t.Run("Packet Event", Parse_Tests(results.PacketEvent, parser.ParsePacket))
	t.Run("Packet FinalClassification", Parse_Tests(results.PacketFinalClassification, parser.ParsePacket))
	t.Run("Packet LapData", Parse_Tests(results.PacketLapData, parser.ParsePacket))
	t.Run("Packet LobbyInfo", Parse_Tests(results.PacketLobbyInfo, parser.ParsePacket))
	t.Run("Packet Motion", Parse_Tests(results.PacketMotion, parser.ParsePacket))
	t.Run("Packet MotionEx", Parse_Tests(results.PacketMotionEx, parser.ParsePacket))
	t.Run("Packet Participants", Parse_Tests(results.PacketParticipants, parser.ParsePacket))
	t.Run("Packet Session", Parse_Tests(results.PacketSession, parser.ParsePacket))
	t.Run("Packet SessionHistory", Parse_Tests(results.PacketSessionHistory, parser.ParsePacket))
	t.Run("Packet TyreSets", Parse_Tests(results.PacketTyreSets, parser.ParsePacket))
}

func Parse_Tests[T any](PacketTyreSets *test_util.DataPoints[T], parse func(decoder *encoding.Decoder) (f1_2023.Packet, error)) func(t *testing.T) {
	return func(t *testing.T) {
		for i, data := range PacketTyreSets.Packets {
			name := fmt.Sprintf("Packet %d", i)
			t.Run(name, func(t *testing.T) {
				decoder := encoding.NewDecoder(data.Data)
				p, err := parse(decoder)

				if data.Error != "" {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					_, ok := p.(T)
					require.True(t, ok)
					// require.Equal(t, data.Packet, &packet)

					header := p.GetHeader()
					require.Equal(t, header.PacketFormat, enums.PF_F1_2023)
					require.GreaterOrEqual(t, header.GameMajorVersion, uint8(1))
					require.GreaterOrEqual(t, header.GameMinorVersion, uint8(20))
					require.Equal(t, header.GameYear, uint8(23))
					require.Equal(t, header.PacketVersion, uint8(1))
				}
			})
		}
	}
}
