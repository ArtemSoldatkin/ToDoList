package todolist

// Sort by name
type byName []Deal

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].Name < b[j].Name }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Sort by description
type byDescription []Deal

func (b byDescription) Len() int           { return len(b) }
func (b byDescription) Less(i, j int) bool { return b[i].Description < b[j].Description }
func (b byDescription) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Sort by description
type byStartDate []Deal

func (b byStartDate) Len() int           { return len(b) }
func (b byStartDate) Less(i, j int) bool { return b[i].StartDate.Sub(b[j].StartDate) < 0 }
func (b byStartDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// Sort by description
type byEndDate []Deal

func (b byEndDate) Len() int           { return len(b) }
func (b byEndDate) Less(i, j int) bool { return b[i].EndDate.Sub(b[j].EndDate) < 0 }
func (b byEndDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
