package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
type resArray[4]struct{
	name string
	num int
}
func getNext(URLpath string) resArray{
	var result resArray
	var err error

	head, tail := ShiftPath(URLpath)

	for i:=0; head != ""; i++ {
		result[i].name = head
		head, tail = ShiftPath(tail)
		result[i].num, err = strconv.Atoi(head)
		if err != nil {
			result[i].num = -1
			fmt.Println("no number:", head)
		}
		head, tail = ShiftPath(tail)
	}
	return result
}

func (a *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("\nServeHTTP(App): req.URL.Path->", req.URL.Path)

	var ep resArray
	fmt.Println(req.URL.Path)
	ep = getNext(req.URL.Path)
	fmt.Println(ep)

	switch  {
	case ep[0].name == "":
		{
			a.HomeHandler.ServeHTTP(res, req)
			fmt.Println("HomeHandler")
			return
		}
	case ep[0].name == "users":
		{
			a.UserHandler.ServeHTTP(res, req)
			fmt.Println("UserHandler")
			head, tail := ShiftPath(req.URL.Path)
			head, tail = ShiftPath(tail)
			head, tail = ShiftPath(tail)
			fmt.Println("head, tail", head, tail)
			return
		}
	case ep[0].name == "items" && ep[1].name == "users" :
		{
			a.ItemUserHandler.ServeHTTP(res,req)
			fmt.Println("items and users")
		}
	case ep[0].name == "items" :
		{
			a.ItemHandler.ServeHTTP(res, req)
			fmt.Println("ItemHandler")
			return
		}
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
		fmt.Println("ErrorHandler")
	}
}
func (h *HomeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("HomeHandler inside")
}
func (h *ItemHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("ItemHandler inside")
}
func (h *ItemUserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("ItemUserHandler inside")
}
func (h *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	head, tail := ShiftPath(req.URL.Path)
	if tail == "/" {
		fmt.Println("list all Users")
		return
	}
	head, tail = ShiftPath(tail)
	userID, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid userUD %s", head), http.StatusBadRequest)
		return
	}
	if tail == "/" {
		switch req.Method {
		case "GET":
			fmt.Println("user:handleGET id=", userID)
			//		h.handleGet(id)
		case "POST":
			fmt.Println("user:handlePOST id=", userID)
			//		h.handlePost(id)
		case "PUT":
			fmt.Println("user:handlePUT id=", userID)
			//		h.handlePUT(id)
		case "PATCH":
			fmt.Println("user:handlePATCH id=", userID)
			//		h.handlePATCH(id)
		case "DELETE":
			fmt.Println("user:handleDELETE id=", userID)
			//		h.handleDelete(id)
		default:
			http.Error(res, "Only GET, POST, PUT, PATCH, DELETE are allowed", http.StatusMethodNotAllowed)
		}
	} else {
		head, tail = ShiftPath(tail)
		switch head {
		case "items":
			{
				head, tail = ShiftPath(tail)
				if head == "" {
					fmt.Println("List all items of user ", userID)
				} else {
					//					fmt.Println("head, tail:", head, tail)
					itemID, err := strconv.Atoi(head)
					if err != nil {
						http.Error(res, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
						return
					}
					switch req.Method {
					case "GET":
						fmt.Println("item:handleGET UID, IID =", userID, itemID)
						//		h.handleGet(id)
					case "PUT":
						fmt.Println("item:handlePUT UID, IID =", userID, itemID)
						//		h.handlePut(id)
					default:
						http.Error(res, "Only GET and PUT are allowed", http.StatusMethodNotAllowed)
					}
				}
			}
		default:
			fmt.Println("not found", head)
		}
	}
}
type UserHandler struct {
}
type ItemHandler struct {
}
type ItemUserHandler struct {
}
type HomeHandler struct {
}
type App struct {
	// We could use http.Handler as a type here; using the specific type has
	// the advantage that static analysis tools can link directly from
	// h.UserHandler.ServeHTTP to the correct definition. The disadvantage is
	// that we have slightly stronger coupling. Do the tradeoff yourself.
	UserHandler *UserHandler
	ItemHandler *ItemHandler
	ItemUserHandler *ItemUserHandler
	HomeHandler *HomeHandler
}
func main() {
	a := &App{
		HomeHandler: new(HomeHandler),
		ItemHandler: new(ItemHandler),
		ItemUserHandler: new(ItemUserHandler),
		UserHandler: new(UserHandler),
	}

	err := http.ListenAndServe(":80", a)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}