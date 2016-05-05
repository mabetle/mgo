package mcore

type KeyValueArray struct {
	cursor int
	rows   []KeyValue
}

// NewKeyValueArrayFromFile
func NewKeyValueArrayFromFile(location, sep string) (*KeyValueArray, error) {
	lines, err := ReadFileLines(location)
	if err != nil {
		return nil, err
	}
	impl := new(KeyValueArray)
	impl.cursor = -1
	impl.rows = GetKeyValueArray(lines, sep)
	return impl, nil
}

func (kva *KeyValueArray) Next() bool {
	//move cursor to next
	kva.cursor = kva.cursor + 1
	if len(kva.rows) > kva.cursor {
		return true
	}
	// move cursor to top
	kva.MoveTop()
	return false
}

func (kva *KeyValueArray) Key() string {
	return kva.rows[kva.cursor].Key
}

func (kva *KeyValueArray) Value() interface{} {
	return kva.rows[kva.cursor].Value
}

func (kva *KeyValueArray) String() string {
	return kva.rows[kva.cursor].String()
}

func (kva *KeyValueArray) Int(dv int) int {
	return kva.rows[kva.cursor].Int(dv)
}

func (kva *KeyValueArray) Bool() bool {
	return kva.rows[kva.cursor].Bool()
}

func (kva *KeyValueArray) Float64(dv float64) float64 {
	return kva.rows[kva.cursor].Float64(dv)
}

func (kva *KeyValueArray) Keys() []string {
	keys := []string{}
	kva.MoveTop()
	for kva.Next() {
		keys = append(keys, kva.Key())
	}
	return keys
}

func (kva *KeyValueArray) IsContainKey(key string) bool {
	return String(key).IsInArray(kva.Keys())
}

func (kva *KeyValueArray) Put(key string, value interface{}) {
	// blank key not accept.
	if key == "" {
		return
	}
	// update exists
	kva.MoveTop()
	for kva.Next() {
		if key == kva.Key() {
			kva.rows[kva.cursor].Value = value
			return
		}
	}
	// insert not exists.
	kva.rows = append(kva.rows, KeyValue{key, value})
}

func (kva *KeyValueArray) KeyString(key string) (string, bool) {
	kva.MoveTop()
	for kva.Next() {
		if key == kva.Key() {
			return kva.String(), true
		}
	}
	return "", false
}

func (kva *KeyValueArray) KeyInt(key string, dv int) (int, bool) {
	kva.MoveTop()
	for kva.Next() {
		if key == kva.Key() {
			return kva.Int(dv), true
		}
	}
	return dv, false
}

func (kva *KeyValueArray) KeyBool(key string) bool {
	kva.MoveTop()
	for kva.Next() {
		if key == kva.Key() {
			return kva.Bool()
		}
	}
	return false
}

func (kva *KeyValueArray) KeyFloat64(key string, dv float64) (float64, bool) {
	kva.MoveTop()
	for kva.Next() {
		if key == kva.Key() {
			return kva.Float64(dv), true
		}
	}
	return dv, false
}

func (kva *KeyValueArray) MoveTop() {
	kva.cursor = -1
}
