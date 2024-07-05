package pair

// Count - Return the size of the OrderedPair
func (o *OrderedPair[KeyType, ValueType]) Count() int {
	return len(o.data)
}
