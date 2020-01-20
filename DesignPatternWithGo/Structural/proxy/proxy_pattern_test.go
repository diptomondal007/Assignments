package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342345)
	for i:=0;i<1000000;i++{
		n:= rand.Int31()
		someDatabase = append(someDatabase, User{ID:n})
	}

	proxy := UserListProxy{
		SomeDatabase:           someDatabase,
		StackCache:             UserList{},
		StackCapacity:          2,
	}

	knownIds := [3]int32{someDatabase[3].ID,someDatabase[4].ID, someDatabase[5].ID}

	t.Run("Find User - Empty cache", func (t *testing.T){
		user, err := proxy.FindUser(knownIds[0])
		if err != nil{
			t.Fatal(err)
		}

		if user.ID != knownIds[0] {
			t.Error("Returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})
}
