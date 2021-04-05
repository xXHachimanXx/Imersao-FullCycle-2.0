package route

type Route struct {
	ID string
	ClientID string
	Positions []Position
}

type Position struct {
	Lat float64
	Long float64
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route ID not informed!")
	}

	f, err := os.Open("detinations/" + r.ID + ".txt")

	if err != nil {
		return err
	}

	defer f.Close()

	
}