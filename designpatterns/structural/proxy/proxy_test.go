package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	mockedDatabase := UserList{}
	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, User{ID: n})
	}
	proxy := UserListProxy{
		MockedDatabase: &mockedDatabase,
		StackSize:      2,
		StackCache:     UserList{},
	}
	knownIDs := [3]int32{mockedDatabase[3].ID, mockedDatabase[4].ID, mockedDatabase[5].ID}
	t.Run("FindUser - Empty Cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesnt match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that  is stored on it.")
		}

		if proxy.LastSearchUsedCache == true {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesnt match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that  is stored on it.")
		}
		if !proxy.LastSearchUsedCache {
			t.Error("the user should have been returned from the cache")
		}
	})

	t.Run("FindUser - Overflowing the stack", func(t *testing.T) {
		user1, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		user2, _ := proxy.FindUser(knownIDs[1])
		if proxy.LastSearchUsedCache {
			t.Error("the user wasnt stored on the cache yet")
		}
		user3, _ := proxy.FindUser(knownIDs[2])
		if proxy.LastSearchUsedCache {
			t.Error("the user wasnt stored on the cache yet")
		}

		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User that should be gone was found")
			}
		}
		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users the cache should not grow to more than 2")
		}

		for _, v := range proxy.StackCache {
			if v != user2 && v != user3 {
				t.Error("Bot users expected to be in cache were not found")
			}
		}
	})
}
