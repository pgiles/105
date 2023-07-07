package enum

// Promotion is enum of either "dynamic" or "elb".
// See Struct-based Enums at https://threedots.tech/post/safer-enums-in-go/
type Promotion struct {
	slug string
}

func (bt *Promotion) String() string {
	return bt.slug
}

var (
	Seasonal  = Promotion{"seasonal"}
	Daily     = Promotion{"daily"}
	Undefined = Promotion{""}
)
