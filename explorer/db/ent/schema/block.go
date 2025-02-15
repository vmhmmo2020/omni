package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Block holds the schema definition for the Block entity.
type Block struct {
	ent.Schema
}

// Fields of the XBlock.
func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("SourceChainID"),
		field.Uint64("BlockHeight"),
		field.Bytes("BlockHash").MaxLen(32),
		field.Time("Timestamp").
			Default(time.Now()),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the Block.
func (Block) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Msgs", Msg.Type),
		edge.To("Receipts", Receipt.Type),
	}
}
