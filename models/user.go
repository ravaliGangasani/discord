package models

import (
	"encoding/json"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	users = "users"
)

// nolint:gochecknoinits
func init() {
	s := session.Copy()
	defer s.Close()

	c := s.DB(info.Database).C(users)

	indexes := []mgo.Index{
		{Key: []string{"id"}, Unique: true},
		{Key: []string{"name"}},
		{Key: []string{"address"}},
		{Key: []string{"claimed"}},
	}

	for i := 0; i < len(indexes); i++ {
		indexes[i].Background = true
		if err := c.EnsureIndex(indexes[i]); err != nil {
			log.Panicln(err)
		}
	}
}

type User struct {
	ID      string `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Claimed bool   `json:"claimed" bson:"claimed"`
}

// func NewNodeFromRaw(n *vpn.Node, ip, port string) *Node {
// 	node := Node{
// 		ID:      n.ID.String(),
// 		Address: common.HexBytes(n.Owner.Bytes()).String(),
// 		Deposit: types.Coin{
// 			Denom: n.Deposit.Denom,
// 			Value: n.Deposit.Amount.Int64(),
// 		},

// 		IP:   ip,
// 		Port: port,

// 		Type:        n.Type,
// 		Version:     n.Version,
// 		Moniker:     n.Moniker,
// 		PricesPerGB: make([]types.Coin, 0, n.PricesPerGB.Len()),
// 		InternetSpeed: types.Bandwidth{
// 			Upload:   n.InternetSpeed.Upload.Int64(),
// 			Download: n.InternetSpeed.Download.Int64(),
// 		},
// 		Encryption: n.Encryption,
// 		Status:     "INACTIVE",
// 	}

// 	for _, c := range n.PricesPerGB {
// 		node.PricesPerGB = append(node.PricesPerGB,
// 			types.Coin{
// 				Denom: c.Denom,
// 				Value: c.Amount.Int64(),
// 			})
// 	}

// 	return &node
// }

func (u *User) Save() error {
	s := session.Copy()
	defer s.Close()

	c := s.DB(info.Database).C(users)
	if err := c.Insert(u); err != nil {
		return err
	}

	return nil
}

func (u *User) String() string {
	bytes, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Panicln(err)
	}

	return string(bytes)
}

func (u *User) MarshalJSON() []byte {
	bytes, err := json.Marshal(u)
	if err != nil {
		log.Panicln(err)
	}

	return bytes
}

func (u *User) FindOne(query bson.M) error {
	s := session.Copy()
	defer s.Close()

	c := s.DB(info.Database).C(users)
	if err := c.Find(query).One(u); err != nil {
		return err
	}

	return nil
}

func (u *User) FindOneAndUpdate(query, update bson.M, upsert, remove, _new bool) error {
	s := session.Copy()
	defer s.Close()

	c := s.DB(info.Database).C(users)
	change := mgo.Change{
		Update: update,
	}

	if _, err := c.Find(query).Apply(change, u); err != nil {
		return err
	}

	return nil
}

type Users []User

func (u *Users) MarshalJSON() []byte {
	bytes, err := json.Marshal(u)
	if err != nil {
		log.Panicln(err)
	}

	return bytes
}

func (u *Users) Find(query, selector bson.M, sort []string, skip, limit int) error {
	s := session.Copy()
	defer s.Close()

	c := s.DB(info.Database).C(users)

	err := c.Find(query).Select(selector).Sort(sort...).Skip(skip).Limit(limit).All(u)
	if err != nil {
		return err
	}

	return nil
}
