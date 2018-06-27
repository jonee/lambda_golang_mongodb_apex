
/**
 * class main
 */

package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/globalsign/mgo"
	// "github.com/globalsign/mgo/bson"

	"crypto/tls"

	// "errors"
	"log"
    "net"
	"os"
    "strconv"
	"time"
)

var t0 = time.Now()
var mongoSession *mgo.Session

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	t1 := time.Now()
	var err error

	log.Println("start handler")
	log.Println("FunctionName, FunctionVersion:", os.Getenv("AWS_LAMBDA_FUNCTION_NAME"), os.Getenv("AWS_LAMBDA_FUNCTION_VERSION"))

	mapContext := make(map[string]interface{})

	mongoDatabase := MONGODB_DATABASE_DEV
	authDatabase := MONGODB_AUTH_DATABASE_DEV
	mongoUser := MONGODB_REGULAR_USER_DEV
	mongoUserPassword := MONGODB_REGULAR_USER_PASSWORD_DEV
	mongoServers := MONGODB_SERVERS_DEV
	mapContext["mongoDatabase"] = mongoDatabase

	t2 := time.Now()

	// connect mongodb if needed or reuse
	if mongoSession == nil {
		tlsConfig := &tls.Config{}

		dialInfo := &mgo.DialInfo{
			Addrs:    mongoServers,
			Database: authDatabase,
			Username: mongoUser,
			Password: mongoUserPassword,
		}
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
		mongoSession, err = mgo.DialWithInfo(dialInfo)

		log.Println("new mongodb connection")

	} else {
		log.Println("re using mongodb")
	}

	mapContext["mongoSession"] = mongoSession
	// defer mongoSession.Close() // don't close!

	t3 := time.Now()

	customerCol := mongoSession.DB(mongoDatabase).C("customer")
	mapContext["customerCol"] = customerCol

    /*
// this is expensive ~1sec on new db so make sure to do it elsewhere
	err = models.CustomerEnsureIndex(mapContext)
	if err != nil {
		log.Println("ERROR", err)
	}

	t4 := time.Now()
    */

	var customerObj Customer
	customerObj.Username = "username_" + strconv.FormatInt(time.Now().Unix(), 10)
	customerObj.CreatedAt = time.Now()

	_, err = customerObj.Save(mapContext)
	if err != nil {
		log.Println("ERROR", err)
	}

	t6 := time.Now()


	// benchmarks before returning
	log.Println("0-1 lambda life", (t1.UnixNano()-t0.UnixNano())/int64(time.Millisecond))
	log.Println("mongo connect", (t3.UnixNano()-t2.UnixNano())/int64(time.Millisecond))

	log.Println("db saving / create customer", (t6.UnixNano()-t3.UnixNano())/int64(time.Millisecond))


    return events.APIGatewayProxyResponse{
        Body:       "Hello ",
        StatusCode: 200,
    }, nil

}

func main() {
	lambda.Start(Handler)
}
