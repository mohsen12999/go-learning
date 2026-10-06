package developer

type Developer struct {
    Name       string
    Role       string
    Experience int
}

func NewDeveloper(name string, role string, experience int) Developer {
    return Developer{
        Name:       name,
        Role:       role,
        Experience: experience,
    }
}

func (developer Developer) IsSenior() bool {
    return developer.Experience >= 7
}
