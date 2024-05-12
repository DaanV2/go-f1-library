package test_util

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/DaanV2/go-f1-library/encoding"
	"github.com/DaanV2/go-f1-library/enums"
	f1_2023 "github.com/DaanV2/go-f1-library/packets/2023"
)

type DataPoint[P any] struct {
	Packet *P     `json:"packet,omitempty"`
	Error  string `json:"error,omitempty"`
	Data   []byte `json:"data,omitempty"`
}

func NewDataPoint[P any](packet *P, err error, data []byte) DataPoint[P] {
	errmsg := ""
	if err != nil {
		errmsg = err.Error()
	}

	return DataPoint[P]{Packet: packet, Error: errmsg, Data: slices.Clone(data)}
}

type DataPoints[P any] struct {
	Packets []DataPoint[P] `json:"packets,omitempty"`
}

func NewDataPoints[P any]() *DataPoints[P] {
	return &DataPoints[P]{Packets: make([]DataPoint[P], 0, 10)}
}

func (d *DataPoints[P]) Add(packet *P, err error, data []byte) {
	if len(d.Packets) > 10 {
		return
	}

	d.Packets = append(d.Packets, NewDataPoint(packet, err, data))
}

func (d *DataPoints[P]) Store(filepath string) error {
	fmt.Println("Storing", filepath)
	// Store packet to file as json
	bytes, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, bytes, 0644)
}

func (d *DataPoints[P]) Load(filepath string) error {
	fmt.Println("Loading", filepath)
	// Load packet from file as json
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, d)
}

type Results struct {
	PacketCarDamage           *DataPoints[f1_2023.PacketCarDamageData]
	PacketCarSetups           *DataPoints[f1_2023.PacketCarSetupsData]
	PacketCarStatus           *DataPoints[f1_2023.PacketCarStatusData]
	PacketCarTelemetry        *DataPoints[f1_2023.PacketCarTelemetryData]
	PacketEvent               *DataPoints[f1_2023.PacketEventData]
	PacketFinalClassification *DataPoints[f1_2023.PacketFinalClassificationData]
	// PacketHeader              *DataPoints[f1_2023.PacketHeader]
	PacketLapData             *DataPoints[f1_2023.PacketLapData]
	PacketLobbyInfo           *DataPoints[f1_2023.PacketLobbyInfoData]
	PacketMotion              *DataPoints[f1_2023.PacketMotionData]
	PacketMotionEx            *DataPoints[f1_2023.PacketMotionExData]
	PacketParticipants        *DataPoints[f1_2023.PacketParticipantsData]
	PacketSession             *DataPoints[f1_2023.PacketSessionData]
	PacketSessionHistory      *DataPoints[f1_2023.PacketSessionHistoryData]
	PacketTyreSets            *DataPoints[f1_2023.PacketTyreSetsData]
}

func NewResults() *Results {
	return &Results{
		PacketCarDamage:           NewDataPoints[f1_2023.PacketCarDamageData](),
		PacketCarSetups:           NewDataPoints[f1_2023.PacketCarSetupsData](),
		PacketCarStatus:           NewDataPoints[f1_2023.PacketCarStatusData](),
		PacketCarTelemetry:        NewDataPoints[f1_2023.PacketCarTelemetryData](),
		PacketEvent:               NewDataPoints[f1_2023.PacketEventData](),
		PacketFinalClassification: NewDataPoints[f1_2023.PacketFinalClassificationData](),
		// PacketHeader:              NewDataPoints[f1_2023.PacketHeader](),
		PacketLapData:             NewDataPoints[f1_2023.PacketLapData](),
		PacketLobbyInfo:           NewDataPoints[f1_2023.PacketLobbyInfoData](),
		PacketMotion:              NewDataPoints[f1_2023.PacketMotionData](),
		PacketMotionEx:            NewDataPoints[f1_2023.PacketMotionExData](),
		PacketParticipants:        NewDataPoints[f1_2023.PacketParticipantsData](),
		PacketSession:             NewDataPoints[f1_2023.PacketSessionData](),
		PacketSessionHistory:      NewDataPoints[f1_2023.PacketSessionHistoryData](),
		PacketTyreSets:            NewDataPoints[f1_2023.PacketTyreSetsData](),
	}
}

func (r *Results) AddPacket(parser *f1_2023.PacketParser, data []byte) {
	// Just header
	decoder := encoding.NewDecoder(data)
	header, herr := parser.PacketHeader(decoder)

	// All data
	decoder = encoding.NewDecoder(data)
	packet, perr := parser.ParsePacket(decoder)
	err := errors.Join(herr, perr)

	// r.PacketHeader.Add(&header, herr, data)

	fmt.Println(" ID:", header.PacketId.String())

	size := parser.Size(packet)
	if size < len(data) {
		data = data[:size]
	}
	data = slices.Clone(data)

	if packet == nil {
		switch header.PacketId {
		case enums.PID_Motion:
			r.PacketMotion.Add(nil, err, data)
		case enums.PID_Session:
			r.PacketSession.Add(nil, err, data)
		case enums.PID_LapData:
			r.PacketLapData.Add(nil, err, data)
		case enums.PID_Event:
			r.PacketEvent.Add(nil, err, data)
		case enums.PID_Participants:
			r.PacketParticipants.Add(nil, err, data)
		case enums.PID_CarSetups:
			r.PacketCarSetups.Add(nil, err, data)
		case enums.PID_CarTelemetry:
			r.PacketCarTelemetry.Add(nil, err, data)
		case enums.PID_CarStatus:
			r.PacketCarStatus.Add(nil, err, data)
		case enums.PID_FinalClassification:
			r.PacketFinalClassification.Add(nil, err, data)
		case enums.PID_LobbyInfo:
			r.PacketLobbyInfo.Add(nil, err, data)
		case enums.PID_CarDamage:
			r.PacketCarDamage.Add(nil, err, data)
		case enums.PID_SessionHistory:
			r.PacketSessionHistory.Add(nil, err, data)
		case enums.PID_TyreSets:
			r.PacketTyreSets.Add(nil, err, data)
		case enums.PID_MotionEx:
			r.PacketMotionEx.Add(nil, err, data)
		}
	}

	switch p := packet.(type) {
	case f1_2023.PacketMotionData:
		r.PacketMotion.Add(&p, err, data)
	case f1_2023.PacketSessionData:
		r.PacketSession.Add(&p, err, data)
	case f1_2023.PacketLapData:
		r.PacketLapData.Add(&p, err, data)
	case f1_2023.PacketEventData:
		r.PacketEvent.Add(&p, err, data)
	case f1_2023.PacketParticipantsData:
		r.PacketParticipants.Add(&p, err, data)
	case f1_2023.PacketCarSetupsData:
		r.PacketCarSetups.Add(&p, err, data)
	case f1_2023.PacketCarTelemetryData:
		r.PacketCarTelemetry.Add(&p, err, data)
	case f1_2023.PacketCarStatusData:
		r.PacketCarStatus.Add(&p, err, data)
	case f1_2023.PacketFinalClassificationData:
		r.PacketFinalClassification.Add(&p, err, data)
	case f1_2023.PacketLobbyInfoData:
		r.PacketLobbyInfo.Add(&p, err, data)
	case f1_2023.PacketCarDamageData:
		r.PacketCarDamage.Add(&p, err, data)
	case f1_2023.PacketSessionHistoryData:
		r.PacketSessionHistory.Add(&p, err, data)
	case f1_2023.PacketTyreSetsData:
		r.PacketTyreSets.Add(&p, err, data)
	case f1_2023.PacketMotionExData:
		r.PacketMotionEx.Add(&p, err, data)
	}
}

func (r *Results) Store(folder string) error {
	if os.MkdirAll(folder, 0755) != nil {
		return errors.New("failed to create folder")
	}

	return errors.Join(
		r.PacketCarDamage.Store(path.Join(folder, "car-damage.json")),
		r.PacketCarSetups.Store(path.Join(folder, "car-setups.json")),
		r.PacketCarStatus.Store(path.Join(folder, "car-status.json")),
		r.PacketCarTelemetry.Store(path.Join(folder, "car-telemetry.json")),
		r.PacketEvent.Store(path.Join(folder, "event.json")),
		r.PacketFinalClassification.Store(path.Join(folder, "final-classification.json")),
		// r.PacketHeader.Store(path.Join(folder, "header.json")),
		r.PacketLapData.Store(path.Join(folder, "lapdata.json")),
		r.PacketLobbyInfo.Store(path.Join(folder, "lobby-info.json")),
		r.PacketMotion.Store(path.Join(folder, "motion.json")),
		r.PacketMotionEx.Store(path.Join(folder, "motion-ex.json")),
		r.PacketParticipants.Store(path.Join(folder, "participants.json")),
		r.PacketSession.Store(path.Join(folder, "session.json")),
		r.PacketSessionHistory.Store(path.Join(folder, "session-history.json")),
		r.PacketTyreSets.Store(path.Join(folder, "tyre-sets.json")),
	)
}

func (r *Results) Load(folder string) error {
	return errors.Join(
		r.PacketCarDamage.Load(path.Join(folder, "car-damage.json")),
		r.PacketCarSetups.Load(path.Join(folder, "car-setups.json")),
		r.PacketCarStatus.Load(path.Join(folder, "car-status.json")),
		r.PacketCarTelemetry.Load(path.Join(folder, "car-telemetry.json")),
		r.PacketEvent.Load(path.Join(folder, "event.json")),
		r.PacketFinalClassification.Load(path.Join(folder, "final-classification.json")),
		// r.PacketHeader.Load(path.Join(folder, "header.json")),
		r.PacketLapData.Load(path.Join(folder, "lapdata.json")),
		r.PacketLobbyInfo.Load(path.Join(folder, "lobby-info.json")),
		r.PacketMotion.Load(path.Join(folder, "motion.json")),
		r.PacketMotionEx.Load(path.Join(folder, "motion-ex.json")),
		r.PacketParticipants.Load(path.Join(folder, "participants.json")),
		r.PacketSession.Load(path.Join(folder, "session.json")),
		r.PacketSessionHistory.Load(path.Join(folder, "session-history.json")),
		r.PacketTyreSets.Load(path.Join(folder, "tyre-sets.json")),
	)
}
