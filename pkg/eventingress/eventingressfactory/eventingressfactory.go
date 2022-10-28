package eventingress

func CreateEventIngressHandler() func() string {
	return func() string {
		return "Ingressed event"
	}
}
