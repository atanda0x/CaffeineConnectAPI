package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var ErrProductNotFound = fmt.Errorf("product not found")

// Produc t defines the structure for an API
type Product struct {
	// the id for this user
	//
	// required: true
	// min: 1
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	Deletedon   string  `json:"-"`
}

// Product is a collection of project
type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-6grt-uu
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().Kind().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

// FromJSON
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON serialized the content of the collection to json
// NewEncoder pro vide better performance than json.unmarshal as it doesn't
// have to buffer the output into an in memory slice of bytes
// this reduce allocation and the overhead of the serv ice
func (p *Products) ToJOSN(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProducts(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[:pos], productList[pos+1])
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// GetProduct return a list of products
func GetProucts() Products {
	return productList
}

// productList is a hard coded list of product for this data store
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "From milky Coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Cappuccino",
		Description: "Espresso-based coffee with steamed milk foam",
		Price:       2.75,
		SKU:         "ghi567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          4,
		Name:        "Macchiato",
		Description: "Espresso 'stained' with a small amount of steamed milk",
		Price:       2.25,
		SKU:         "klm890",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          5,
		Name:        "Americano",
		Description: "Espresso with hot water, resembling American-style coffee",
		Price:       1.75,
		SKU:         "nop234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          6,
		Name:        "Mocha",
		Description: "Espresso with steamed milk and chocolate syrup",
		Price:       2.95,
		SKU:         "qrs456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          7,
		Name:        "Flat White",
		Description: "Espresso-based coffee with velvety microfoam",
		Price:       2.65,
		SKU:         "tuv678",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          8,
		Name:        "Affogato",
		Description: "Espresso poured over a scoop of vanilla ice cream",
		Price:       3.25,
		SKU:         "wxy901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          9,
		Name:        "Turkish Coffee",
		Description: "Unfiltered coffee brewed with finely ground coffee beans",
		Price:       2.55,
		SKU:         "zab345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          10,
		Name:        "Iced Coffee",
		Description: "Chilled coffee served over ice",
		Price:       2.35,
		SKU:         "cde567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          11,
		Name:        "Croissant",
		Description: "Buttery, flaky pastry",
		Price:       1.75,
		SKU:         "fgh789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          12,
		Name:        "Bagel",
		Description: "Dense, chewy bread roll",
		Price:       1.25,
		SKU:         "hij901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          13,
		Name:        "Muffin",
		Description: "Individual-sized quick bread product",
		Price:       1.50,
		SKU:         "jkl234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          14,
		Name:        "Drip Coffee",
		Description: "Brewed coffee made by pouring hot water over ground coffee",
		Price:       1.65,
		SKU:         "mno456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          15,
		Name:        "Tea",
		Description: "Hot or cold infusion of dried herbs, flowers, or leaves",
		Price:       1.25,
		SKU:         "pqr678",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          16,
		Name:        "Hot Chocolate",
		Description: "Hot drink made with chocolate and milk",
		Price:       2.25,
		SKU:         "stu890",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          17,
		Name:        "Fruit Smoothie",
		Description: "Blended drink made with fruits, yogurt, and ice",
		Price:       3.25,
		SKU:         "vwx123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          18,
		Name:        "Iced Tea",
		Description: "Chilled tea served over ice",
		Price:       1.75,
		SKU:         "yza345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          19,
		Name:        "Pastry",
		Description: "Sweet baked goods made from dough",
		Price:       2.45,
		SKU:         "bcd567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          20,
		Name:        "Sandwich",
		Description: "Food typically consisting of vegetables, sliced cheese or meat, placed on or between slices of bread",
		Price:       3.75,
		SKU:         "cde789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          21,
		Name:        "Cookie",
		Description: "Sweet baked treat",
		Price:       1.25,
		SKU:         "efg901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          22,
		Name:        "Brownie",
		Description: "Rich, chocolate dessert",
		Price:       1.75,
		SKU:         "ghi123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          23,
		Name:        "Scone",
		Description: "Single-serving cake or quick bread",
		Price:       1.50,
		SKU:         "ijk345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          24,
		Name:        "Smoothie",
		Description: "Thick, smooth drink made from pureed raw fruit, vegetables, or ice cream",
		Price:       3.25,
		SKU:         "klm567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          25,
		Name:        "Fruit Juice",
		Description: "Drink made from the extraction or pressing of the natural liquid contained in fruit",
		Price:       2.00,
		SKU:         "mno789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          26,
		Name:        "Baguette",
		Description: "Long, thin loaf of French bread",
		Price:       2.25,
		SKU:         "pqr901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          27,
		Name:        "Pancake",
		Description: "Flat cake, often thin and round, prepared from a starch-based batter",
		Price:       2.50,
		SKU:         "stu123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          28,
		Name:        "Waffle",
		Description: "Batter-based or dough-based cake cooked in a waffle iron patterned to give a characteristic size, shape, and surface impression",
		Price:       2.75,
		SKU:         "vwx345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          29,
		Name:        "French Toast",
		Description: "Sliced bread soaked in beaten eggs and typically milk, then pan-fried",
		Price:       2.75,
		SKU:         "yza567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          30,
		Name:        "Omelette",
		Description: "Egg dish made from beaten eggs quickly cooked with butter or oil in a frying pan",
		Price:       3.00,
		SKU:         "bcd789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          31,
		Name:        "Quiche",
		Description: "Savory, open-faced pastry crust with a filling of savory custard with cheese, meat, seafood, and/or vegetables",
		Price:       3.25,
		SKU:         "def901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          32,
		Name:        "Salad",
		Description: "Mixture of small pieces of food, usually vegetables or fruits",
		Price:       3.50,
		SKU:         "efg123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          33,
		Name:        "Pasta",
		Description: "Italian food typically made from an unleavened dough of wheat flour mixed with water or eggs, and formed into sheets or various shapes, then cooked by boiling or baking",
		Price:       3.75,
		SKU:         "fgh345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          34,
		Name:        "Pizza",
		Description: "Italian dish consisting of a usually round, flattened base of leavened wheat-based dough topped with tomatoes, cheese, and often various other ingredients",
		Price:       4.00,
		SKU:         "ghi567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          35,
		Name:        "Burger",
		Description: "Sandwich consisting of one or more cooked patties of ground meat, usually beef, placed inside a sliced bread roll or bun",
		Price:       4.25,
		SKU:         "hij789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          36,
		Name:        "Taco",
		Description: "Traditional Mexican dish consisting of a small hand-sized corn or wheat tortilla topped with a filling",
		Price:       3.25,
		SKU:         "ijk901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          37,
		Name:        "Sushi",
		Description: "Japanese dish consisting of small balls or rolls of vinegar-flavored cold cooked rice served with a garnish of raw fish, vegetables, or egg",
		Price:       4.50,
		SKU:         "klm123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          38,
		Name:        "Sashimi",
		Description: "Japanese delicacy consisting of very fresh raw fish or meat sliced into thin pieces",
		Price:       4.75,
		SKU:         "mno345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          39,
		Name:        "Ramen",
		Description: "Japanese noodle soup dish",
		Price:       3.75,
		SKU:         "nop567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          40,
		Name:        "Pho",
		Description: "Vietnamese soup consisting of broth, rice noodles, herbs, and meat",
		Price:       4.25,
		SKU:         "opq789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          41,
		Name:        "Pad Thai",
		Description: "Stir-fried rice noodle dish commonly served as a street food and at most restaurants in Thailand",
		Price:       5.00,
		SKU:         "pqr901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          42,
		Name:        "Hamburger",
		Description: "Sandwich consisting of one or more cooked patties of ground meat, usually beef, placed inside a sliced bread roll or bun",
		Price:       4.50,
		SKU:         "qrs123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          43,
		Name:        "Cheeseburger",
		Description: "Hamburger topped with cheese",
		Price:       5.00,
		SKU:         "rst345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          44,
		Name:        "Banh Mi",
		Description: "Vietnamese sandwich",
		Price:       3.75,
		SKU:         "stu567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          45,
		Name:        "Pho",
		Description: "Vietnamese soup consisting of broth, rice noodles, herbs, and meat",
		Price:       4.25,
		SKU:         "tuv789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          46,
		Name:        "Sushi",
		Description: "Japanese dish consisting of small balls or rolls of vinegar-flavored cold cooked rice served with a garnish of raw fish, vegetables, or egg",
		Price:       4.50,
		SKU:         "uvw901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          47,
		Name:        "Sashimi",
		Description: "Japanese delicacy consisting of very fresh raw fish or meat sliced into thin pieces",
		Price:       4.75,
		SKU:         "wxy123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          48,
		Name:        "Ramen",
		Description: "Japanese noodle soup dish",
		Price:       3.75,
		SKU:         "xyz345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          49,
		Name:        "Gyoza",
		Description: "Japanese dumplings",
		Price:       3.25,
		SKU:         "abc567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          50,
		Name:        "Tempura",
		Description: "Japanese dish usually consisting of seafood or vegetables that have been battered and deep-fried",
		Price:       4.25,
		SKU:         "def789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          51,
		Name:        "Okonomiyaki",
		Description: "Japanese savory pancake containing a variety of ingredients",
		Price:       4.00,
		SKU:         "ghi901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          52,
		Name:        "Takoyaki",
		Description: "Japanese snack in the shape of little round balls containing pieces of octopus",
		Price:       3.50,
		SKU:         "jkl123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          53,
		Name:        "Yakitori",
		Description: "Japanese skewered chicken",
		Price:       3.75,
		SKU:         "mno345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          54,
		Name:        "Onigiri",
		Description: "Japanese rice ball",
		Price:       2.25,
		SKU:         "pqr567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          55,
		Name:        "Miso Soup",
		Description: "Japanese soup consisting of a stock called 'dashi' into which softened miso paste is mixed",
		Price:       2.75,
		SKU:         "stu789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          56,
		Name:        "Tonkatsu",
		Description: "Japanese dish which consists of a breaded, deep-fried pork cutlet",
		Price:       5.25,
		SKU:         "vwx901",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          57,
		Name:        "Karaage",
		Description: "Japanese cooking technique in which various foods—most often chicken, but also other meat and fish—are deep fried in oil",
		Price:       4.75,
		SKU:         "yza123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          58,
		Name:        "Soba",
		Description: "Thin noodles made from buckwheat flour",
		Price:       3.25,
		SKU:         "bcd345",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          59,
		Name:        "Udon",
		Description: "Thick wheat flour noodles",
		Price:       3.50,
		SKU:         "cde567",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          60,
		Name:        "Matcha",
		Description: "Japanese green tea powder",
		Price:       2.75,
		SKU:         "efg789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          61,
		Name:        "Ewa Agoyin",
		Description: "Mashed beans served with a spicy pepper sauce",
		Price:       4.00,
		SKU:         "ewa_agoyin_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          62,
		Name:        "Suya",
		Description: "Spicy skewered meat grilled over an open flame",
		Price:       5.50,
		SKU:         "suya_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          63,
		Name:        "Moin Moin",
		Description: "Steamed bean pudding made from grounded peeled black-eyed peas",
		Price:       3.00,
		SKU:         "moin_moin_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          64,
		Name:        "Ofada Rice and Ayamase",
		Description: "Local Nigerian rice served with a spicy stew made with green bell peppers and assorted meats",
		Price:       7.00,
		SKU:         "ofada_rice_ayamase_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          65,
		Name:        "Amala and Ewedu Soup",
		Description: "Amala is a thick brown paste made from yam flour, served with a soup made from jute leaves",
		Price:       6.50,
		SKU:         "amala_ewedu_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          66,
		Name:        "Abacha",
		Description: "A salad made from dried shredded cassava, vegetables, and palm oil",
		Price:       4.50,
		SKU:         "abacha_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          67,
		Name:        "Banga Soup and Starch",
		Description: "A rich palm nut soup served with a starchy side dish made from cassava",
		Price:       7.50,
		SKU:         "banga_soup_starch_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          68,
		Name:        "Edikang Ikong Soup",
		Description: "A nutritious soup made from a mixture of vegetables, meat, and fish",
		Price:       8.00,
		SKU:         "edikang_ikong_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          69,
		Name:        "Gbure",
		Description: "A traditional Nigerian soup made from okra",
		Price:       5.00,
		SKU:         "gbure_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          70,
		Name:        "Ikokore",
		Description: "A yam-based pottage dish popular among the Ijebu people of Nigeria",
		Price:       6.00,
		SKU:         "ikokore_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          71,
		Name:        "Ogbono Soup",
		Description: "A thick, hearty soup made from ground ogbono seeds",
		Price:       5.50,
		SKU:         "ogbono_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          72,
		Name:        "Ofe Akwu",
		Description: "A delicious soup made from palm fruit extract, popular in the eastern part of Nigeria",
		Price:       7.00,
		SKU:         "ofe_akwu_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          73,
		Name:        "Oha Soup",
		Description: "A traditional soup made from Ora leaves and thickened with cocoyam paste",
		Price:       6.50,
		SKU:         "oha_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          74,
		Name:        "Owo Soup",
		Description: "A delicious soup made from palm nuts and assorted meats and fish",
		Price:       7.50,
		SKU:         "owo_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          75,
		Name:        "Pepper Soup",
		Description: "A spicy soup made with assorted meats and fish, flavored with a blend of spices",
		Price:       6.00,
		SKU:         "pepper_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          76,
		Name:        "Ukwa",
		Description: "A tasty dish made from African breadfruit seeds",
		Price:       6.00,
		SKU:         "ukwa_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          77,
		Name:        "Agidi",
		Description: "A solidified jelly-like pudding made from corn flour",
		Price:       3.00,
		SKU:         "agidi_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          78,
		Name:        "Egusi Soup",
		Description: "A rich soup made from melon seeds, vegetables, and meat or fish",
		Price:       6.50,
		SKU:         "egusi_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          79,
		Name:        "Igbo Bitterleaf Soup",
		Description: "A traditional soup made from bitter leaves and assorted meats",
		Price:       7.00,
		SKU:         "igbo_bitterleaf_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          80,
		Name:        "Isi Ewu",
		Description: "A traditional Igbo soup made from goat head meat",
		Price:       8.00,
		SKU:         "isi_ewu_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          81,
		Name:        "Nkwobi",
		Description: "Spicy cow foot or cow leg dish native to the Igbo people of Nigeria",
		Price:       7.50,
		SKU:         "nkwobi_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          82,
		Name:        "Ofe Nsala",
		Description: "A traditional Igbo soup made with catfish or other types of meat",
		Price:       8.00,
		SKU:         "ofe_nsala_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          83,
		Name:        "Ugba",
		Description: "A salad made from shredded oil bean seeds",
		Price:       5.50,
		SKU:         "ugba_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          84,
		Name:        "Abak Atama",
		Description: "A traditional soup made from palm fruit extract and Atama leaves",
		Price:       7.00,
		SKU:         "abak_atama_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          85,
		Name:        "Afia Efere",
		Description: "A traditional soup made from palm fruit extract and fish",
		Price:       6.50,
		SKU:         "afia_efere_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          86,
		Name:        "Efere Nkpa",
		Description: "A traditional soup made from fresh or dried catfish",
		Price:       6.00,
		SKU:         "efere_nkpa_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          87,
		Name:        "Efere Okazi",
		Description: "A traditional soup made from Okazi leaves and thickened with cocoyam paste",
		Price:       6.50,
		SKU:         "efere_okazi_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          88,
		Name:        "Igbo Ofe Oha",
		Description: "A traditional soup made from Ora leaves and thickened with cocoyam paste",
		Price:       6.50,
		SKU:         "igbo_ofe_oha_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          89,
		Name:        "Ikwerre Banga",
		Description: "A traditional soup made from palm fruit extract and assorted meats",
		Price:       7.50,
		SKU:         "ikwerre_banga_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          90,
		Name:        "Ikwerre Ofe Akwu",
		Description: "A traditional soup made from palm fruit extract, popular in the eastern part of Nigeria",
		Price:       7.00,
		SKU:         "ikwerre_ofe_akwu_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          91,
		Name:        "Ikwere Onugbu",
		Description: "A traditional soup made from bitter leaves and assorted meats",
		Price:       7.00,
		SKU:         "ikwere_onugbu_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          92,
		Name:        "Ikwerre Ukazi",
		Description: "A traditional soup made from Okazi leaves and thickened with cocoyam paste",
		Price:       6.50,
		SKU:         "ikwerre_ukazi_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          93,
		Name:        "Ikwere Ofe Owerri",
		Description: "A traditional soup made from a variety of vegetables and assorted meats",
		Price:       7.50,
		SKU:         "ikwere_ofe_owerri_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          94,
		Name:        "Nkoyo Soup",
		Description: "A traditional soup made from fresh or dried catfish",
		Price:       6.00,
		SKU:         "nkoyo_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          95,
		Name:        "Nsala Soup",
		Description: "A traditional soup made from catfish or other types of meat",
		Price:       6.50,
		SKU:         "nsala_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          96,
		Name:        "Ofe Oha",
		Description: "A traditional soup made from Ora leaves and thickened with cocoyam paste",
		Price:       6.50,
		SKU:         "ofe_oha_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          97,
		Name:        "Ofe Owerri",
		Description: "A traditional soup made from a variety of vegetables and assorted meats",
		Price:       7.50,
		SKU:         "ofe_owerri_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          98,
		Name:        "Ofe Nsala",
		Description: "A traditional soup made with catfish or other types of meat",
		Price:       8.00,
		SKU:         "ofe_nsala_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          99,
		Name:        "Ofe Akwu",
		Description: "A delicious soup made from palm fruit extract, popular in the eastern part of Nigeria",
		Price:       7.00,
		SKU:         "ofe_akwu_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          100,
		Name:        "Owo Soup",
		Description: "A delicious soup made from palm nuts and assorted meats and fish",
		Price:       7.50,
		SKU:         "owo_soup_001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
