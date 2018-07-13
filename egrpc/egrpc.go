package egrpc

// Name contains a first and family name.
type Name struct {
	First string
	Last  string
}

// Email contains the type and actual email address.
type Email struct {
	Type    string
	Address string
}

// Person is an object that holds the name and email of an individual.
type Person struct {
	Name  Name
	Email []Email
}

// PersonData is a type that holds a person's data in a one-line format.
type PersonData struct {
	Data string
}

// FirstLast is an RPC method that displays a person's name in `First Last` format.
func (p *Person) FirstLast(args *Person, reply *string) error {
	*reply = args.Name.First + " " + args.Name.Last
	return nil
}

// LastFirst is an RPC method that displays a person's name in `Last, First` format.
func (p *Person) LastFirst(args *Person, bp *PersonData) error {
	bp.Data = args.Name.Last + ", " + args.Name.First
	return nil
}

// Bio is an RPC method that displays a person's biographic data.
func (p *Person) Bio(args *Person, reply *string) error {
	*reply = "\tName: " + args.Name.First + " " + args.Name.Last
	for _, v := range args.Email {
		*reply += "\n\t\t" + v.Type + ": " + v.Address
	}
	*reply += "\n"
	return nil
}

// SampleData is an array of person objects.
var SampleData = []Person{
	Person{
		Name: Name{
			Last:  "Ritchie",
			First: "Dennis",
		},
		Email: []Email{
			{
				Type:    "work",
				Address: "dmr@unix-labs.com.com",
			},
		},
	},
	Person{
		Name: Name{
			Last:  "Thompson",
			First: "Ken",
		},
		Email: []Email{
			{
				Type:    "work",
				Address: "ken@unix-labs.com.com",
			},
		},
	},
	Person{
		Name: Name{
			Last:  "Pike",
			First: "Rob",
		},
		Email: []Email{
			{
				Type:    "work",
				Address: "rob@unix-labs.com",
			},
			{
				Type:    "home",
				Address: "rob@gmail.com",
			},
		},
	},
}
