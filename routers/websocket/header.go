package router

const (
	HeaderSession string = "Session"
	HeaderToken          = "Token"
)

type Header struct {
	M map[string]interface{}
}

func NewHeader() *Header {
	h := &Header{
		M: make(map[string]interface{}),
	}
	h.Set(HeaderSession, "")
	h.Set(HeaderToken, "")
	return h
}

func (h *Header) Set(key string, value interface{}) {
	if len(key) != 0 {
		h.M[key] = value
	}
}

func (h *Header) Get(key string) interface{} {
	if len(key) == 0 {
		return nil
	} else if v, ok := h.M[key]; ok {
		return v
	}
	return nil
}

func (h *Header) GetString(key string) string {
	if value := h.Get(key); value == nil {
	} else if str, ok := value.(string); ok {
		return str
	}
	return ""
}

func (h *Header) Session() string {
	return h.GetString(HeaderSession)
}

func (h *Header) Token() string {
	return h.GetString(HeaderToken)
}
