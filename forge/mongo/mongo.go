package mongo

//TODO: We can and should create a lib with this boilerplate code

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

//TODO: Check all types in here and try to reorder them as needed

// Client represents a MongoDB client
type Client struct {
	Database *mgo.Database
	session  *mgo.Session
	db       string
}

//Index represents a MGO index
type Index struct {
	Collection string
	Keys       []string
	Unique     bool
}

//SessionConf represents the basic configuration needed for setting up a session
type SessionConf struct {
	//TODO: We could try and use DialInfo to extend
	MongoHosts []string
	DB         string
	User       string
	Password   string
	Indexes    []Index
}

// Connect connects to MongoDB and returns a client
func Connect(sessionConf SessionConf) (*Client, error) {
	log.Infof("Connecting to MongoDB @ %s", sessionConf.MongoHosts)

	dialInfo := &mgo.DialInfo{
		Addrs:    sessionConf.MongoHosts,
		Database: sessionConf.DB,
		Username: sessionConf.User,
		Password: sessionConf.Password,
		Timeout:  time.Second * 10,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	createIndices(session, sessionConf)
	database := session.DB(sessionConf.DB)
	return &Client{
		Database: database,
		session:  session,
	}, nil
}

func createIndices(s *mgo.Session, sessionConf SessionConf) {
	session := s.Copy()
	defer session.Close()
	for _, i := range sessionConf.Indexes {
		c := session.DB(sessionConf.DB).C(i.Collection)
		createIndex(c, i.Keys, i.Unique)
	}

}

func createIndex(c *mgo.Collection, keys []string, unique bool) {
	i := mgo.Index{
		Key:        keys,
		Unique:     unique,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(i)
	if err != nil {
		log.Fatal(err)
	}
}

// Copy creates a new session by calling session.Copy
func (c Client) Copy() Client {
	session := c.session.Copy()
	database := session.DB(c.db)
	return Client{
		Database: database,
		session:  session,
		db:       c.db,
	}
}

// Close closes the MongoDB session
func (c Client) Close() {
	c.session.Close()
}
