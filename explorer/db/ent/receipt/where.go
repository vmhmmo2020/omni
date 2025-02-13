// Code generated by ent, DO NOT EDIT.

package receipt

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "UUID" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldUUID, v))
}

// GasUsed applies equality check predicate on the "GasUsed" field. It's identical to GasUsedEQ.
func GasUsed(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldGasUsed, v))
}

// Success applies equality check predicate on the "Success" field. It's identical to SuccessEQ.
func Success(v bool) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldSuccess, v))
}

// RelayerAddress applies equality check predicate on the "RelayerAddress" field. It's identical to RelayerAddressEQ.
func RelayerAddress(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldRelayerAddress, v))
}

// SourceChainID applies equality check predicate on the "SourceChainID" field. It's identical to SourceChainIDEQ.
func SourceChainID(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldSourceChainID, v))
}

// DestChainID applies equality check predicate on the "DestChainID" field. It's identical to DestChainIDEQ.
func DestChainID(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldDestChainID, v))
}

// StreamOffset applies equality check predicate on the "StreamOffset" field. It's identical to StreamOffsetEQ.
func StreamOffset(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldStreamOffset, v))
}

// TxHash applies equality check predicate on the "TxHash" field. It's identical to TxHashEQ.
func TxHash(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldTxHash, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldCreatedAt, v))
}

// UUIDEQ applies the EQ predicate on the "UUID" field.
func UUIDEQ(v uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "UUID" field.
func UUIDNEQ(v uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "UUID" field.
func UUIDIn(vs ...uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "UUID" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "UUID" field.
func UUIDGT(v uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "UUID" field.
func UUIDGTE(v uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "UUID" field.
func UUIDLT(v uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "UUID" field.
func UUIDLTE(v uuid.UUID) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldUUID, v))
}

// GasUsedEQ applies the EQ predicate on the "GasUsed" field.
func GasUsedEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldGasUsed, v))
}

// GasUsedNEQ applies the NEQ predicate on the "GasUsed" field.
func GasUsedNEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldGasUsed, v))
}

// GasUsedIn applies the In predicate on the "GasUsed" field.
func GasUsedIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldGasUsed, vs...))
}

// GasUsedNotIn applies the NotIn predicate on the "GasUsed" field.
func GasUsedNotIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldGasUsed, vs...))
}

// GasUsedGT applies the GT predicate on the "GasUsed" field.
func GasUsedGT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldGasUsed, v))
}

// GasUsedGTE applies the GTE predicate on the "GasUsed" field.
func GasUsedGTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldGasUsed, v))
}

// GasUsedLT applies the LT predicate on the "GasUsed" field.
func GasUsedLT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldGasUsed, v))
}

// GasUsedLTE applies the LTE predicate on the "GasUsed" field.
func GasUsedLTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldGasUsed, v))
}

// SuccessEQ applies the EQ predicate on the "Success" field.
func SuccessEQ(v bool) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldSuccess, v))
}

// SuccessNEQ applies the NEQ predicate on the "Success" field.
func SuccessNEQ(v bool) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldSuccess, v))
}

// RelayerAddressEQ applies the EQ predicate on the "RelayerAddress" field.
func RelayerAddressEQ(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldRelayerAddress, v))
}

// RelayerAddressNEQ applies the NEQ predicate on the "RelayerAddress" field.
func RelayerAddressNEQ(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldRelayerAddress, v))
}

// RelayerAddressIn applies the In predicate on the "RelayerAddress" field.
func RelayerAddressIn(vs ...[]byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldRelayerAddress, vs...))
}

// RelayerAddressNotIn applies the NotIn predicate on the "RelayerAddress" field.
func RelayerAddressNotIn(vs ...[]byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldRelayerAddress, vs...))
}

// RelayerAddressGT applies the GT predicate on the "RelayerAddress" field.
func RelayerAddressGT(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldRelayerAddress, v))
}

// RelayerAddressGTE applies the GTE predicate on the "RelayerAddress" field.
func RelayerAddressGTE(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldRelayerAddress, v))
}

// RelayerAddressLT applies the LT predicate on the "RelayerAddress" field.
func RelayerAddressLT(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldRelayerAddress, v))
}

// RelayerAddressLTE applies the LTE predicate on the "RelayerAddress" field.
func RelayerAddressLTE(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldRelayerAddress, v))
}

// SourceChainIDEQ applies the EQ predicate on the "SourceChainID" field.
func SourceChainIDEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldSourceChainID, v))
}

// SourceChainIDNEQ applies the NEQ predicate on the "SourceChainID" field.
func SourceChainIDNEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldSourceChainID, v))
}

// SourceChainIDIn applies the In predicate on the "SourceChainID" field.
func SourceChainIDIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldSourceChainID, vs...))
}

// SourceChainIDNotIn applies the NotIn predicate on the "SourceChainID" field.
func SourceChainIDNotIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldSourceChainID, vs...))
}

// SourceChainIDGT applies the GT predicate on the "SourceChainID" field.
func SourceChainIDGT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldSourceChainID, v))
}

// SourceChainIDGTE applies the GTE predicate on the "SourceChainID" field.
func SourceChainIDGTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldSourceChainID, v))
}

// SourceChainIDLT applies the LT predicate on the "SourceChainID" field.
func SourceChainIDLT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldSourceChainID, v))
}

// SourceChainIDLTE applies the LTE predicate on the "SourceChainID" field.
func SourceChainIDLTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldSourceChainID, v))
}

// DestChainIDEQ applies the EQ predicate on the "DestChainID" field.
func DestChainIDEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldDestChainID, v))
}

// DestChainIDNEQ applies the NEQ predicate on the "DestChainID" field.
func DestChainIDNEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldDestChainID, v))
}

// DestChainIDIn applies the In predicate on the "DestChainID" field.
func DestChainIDIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldDestChainID, vs...))
}

// DestChainIDNotIn applies the NotIn predicate on the "DestChainID" field.
func DestChainIDNotIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldDestChainID, vs...))
}

// DestChainIDGT applies the GT predicate on the "DestChainID" field.
func DestChainIDGT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldDestChainID, v))
}

// DestChainIDGTE applies the GTE predicate on the "DestChainID" field.
func DestChainIDGTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldDestChainID, v))
}

// DestChainIDLT applies the LT predicate on the "DestChainID" field.
func DestChainIDLT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldDestChainID, v))
}

// DestChainIDLTE applies the LTE predicate on the "DestChainID" field.
func DestChainIDLTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldDestChainID, v))
}

// StreamOffsetEQ applies the EQ predicate on the "StreamOffset" field.
func StreamOffsetEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldStreamOffset, v))
}

// StreamOffsetNEQ applies the NEQ predicate on the "StreamOffset" field.
func StreamOffsetNEQ(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldStreamOffset, v))
}

// StreamOffsetIn applies the In predicate on the "StreamOffset" field.
func StreamOffsetIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldStreamOffset, vs...))
}

// StreamOffsetNotIn applies the NotIn predicate on the "StreamOffset" field.
func StreamOffsetNotIn(vs ...uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldStreamOffset, vs...))
}

// StreamOffsetGT applies the GT predicate on the "StreamOffset" field.
func StreamOffsetGT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldStreamOffset, v))
}

// StreamOffsetGTE applies the GTE predicate on the "StreamOffset" field.
func StreamOffsetGTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldStreamOffset, v))
}

// StreamOffsetLT applies the LT predicate on the "StreamOffset" field.
func StreamOffsetLT(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldStreamOffset, v))
}

// StreamOffsetLTE applies the LTE predicate on the "StreamOffset" field.
func StreamOffsetLTE(v uint64) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldStreamOffset, v))
}

// TxHashEQ applies the EQ predicate on the "TxHash" field.
func TxHashEQ(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldTxHash, v))
}

// TxHashNEQ applies the NEQ predicate on the "TxHash" field.
func TxHashNEQ(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldTxHash, v))
}

// TxHashIn applies the In predicate on the "TxHash" field.
func TxHashIn(vs ...[]byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldTxHash, vs...))
}

// TxHashNotIn applies the NotIn predicate on the "TxHash" field.
func TxHashNotIn(vs ...[]byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldTxHash, vs...))
}

// TxHashGT applies the GT predicate on the "TxHash" field.
func TxHashGT(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldTxHash, v))
}

// TxHashGTE applies the GTE predicate on the "TxHash" field.
func TxHashGTE(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldTxHash, v))
}

// TxHashLT applies the LT predicate on the "TxHash" field.
func TxHashLT(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldTxHash, v))
}

// TxHashLTE applies the LTE predicate on the "TxHash" field.
func TxHashLTE(v []byte) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldTxHash, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.Receipt {
	return predicate.Receipt(sql.FieldLTE(FieldCreatedAt, v))
}

// HasBlock applies the HasEdge predicate on the "Block" edge.
func HasBlock() predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BlockTable, BlockColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockWith applies the HasEdge predicate on the "Block" edge with a given conditions (other predicates).
func HasBlockWith(preds ...predicate.Block) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := newBlockStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMsgs applies the HasEdge predicate on the "Msgs" edge.
func HasMsgs() predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MsgsTable, MsgsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMsgsWith applies the HasEdge predicate on the "Msgs" edge with a given conditions (other predicates).
func HasMsgsWith(preds ...predicate.Msg) predicate.Receipt {
	return predicate.Receipt(func(s *sql.Selector) {
		step := newMsgsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(sql.NotPredicates(p))
}
