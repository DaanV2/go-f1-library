package f1_2023_test

import (
	"encoding/binary"
	"testing"

	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
	f1_2023 "github.com/DaanV2/go-f1-library/packets/2023"
	"github.com/stretchr/testify/require"
)

func Test_DeserializePacketHeader(t *testing.T) {
	data := make([]byte, 29)

	binary.LittleEndian.PutUint16(data[0:2], uint16(enums.PF_F1_2023))
	data[2] = 23                                  // Game year
	data[3] = 1                                   // Game major version
	data[4] = 0                                   // Game version
	data[5] = 1                                   // Packet version
	data[6] = 0                                   // Packet id
	binary.LittleEndian.PutUint64(data[7:15], 0)  // Session UID
	binary.LittleEndian.PutUint32(data[15:19], 0) // Session time
	binary.LittleEndian.PutUint32(data[19:23], 0) // Frame identifier
	binary.LittleEndian.PutUint32(data[23:27], 0) // Overall frame identifier
	data[27] = 0                                  // Player car index
	data[28] = 255                                // Secondary player car index

	decoder := encoding.NewDecoder(data)
	packet, err := f1_2023.ParsePacketHeader(decoder)
	require.NoError(t, err)

	require.Equal(t, enums.PF_F1_2023, packet.PacketFormat)
	require.Equal(t, uint8(23), packet.GameYear)
	require.Equal(t, uint8(1), packet.GameMajorVersion)
	require.Equal(t, uint8(0), packet.GameMinorVersion)
	require.Equal(t, uint8(1), packet.PacketVersion)
	require.Equal(t, enums.PID_Motion, packet.PacketId)
	require.Equal(t, uint64(0), packet.SessionUID)
	require.Equal(t, float32(0), packet.SessionTime)
	require.Equal(t, uint32(0), packet.FrameIdentifier)
	require.Equal(t, uint32(0), packet.OverallFrameIdentifier)
	require.Equal(t, uint8(0), packet.PlayerCarIndex)
	require.Equal(t, uint8(255), packet.SecondaryPlayerCarIndex)

	require.False(t, packet.HasSecondaryPlayer())
}
