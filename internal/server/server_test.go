package server

	
import (
	"net/http"
	"net/http/httptest"
	"testing"
    "fmt"
    "errors"
    "github.com/massicer/players/internal/store"
)

type StubPlayerStore struct {
    scores map[string]int
    WinCalls []string
    returnErrorDuringStore bool
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
    score := s.scores[name]
    return score
}

func (s *StubPlayerStore) RecordWin(name string)  error {
    if !s.returnErrorDuringStore{
        s.WinCalls = append(s.WinCalls, name)
        return nil
    }
    return errors.New("Cannot insert record")
}

func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
        map[string]int{
            "Pepper": 20,
            "Floyd":  10,
        },
        make([]string, 0),
        false,
    }
    server := &PlayerServer{Store: &store}

    t.Run("returns Pepper's score", func(t *testing.T) {
        request := newGetScoreRequest("Pepper")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertStatus(t, response.Code, http.StatusOK)
    })

    t.Run("returns Floyd's score", func(t *testing.T) {
        request := newGetScoreRequest("Floyd")
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertStatus(t, response.Code, http.StatusOK)
	})
	
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()
	
		server.ServeHTTP(response, request)
	
		got := response.Code
		want := http.StatusNotFound
	
		assertStatus(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
    return req
}

func assertResponseBody(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("response body is wrong, got %q want %q", got, want)
    }
}

func assertStatus(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("did not get correct status, got %d, want %d", got, want)
    }
}

func TestStoreWins(t *testing.T) {
    store := StubPlayerStore{
        map[string]int{},
        make([]string, 0),
        true,
    }
    server := &PlayerServer{&store}

    t.Run("it doesn't record wins when POST if repo returns error", func(t *testing.T) {
        player_name := "Max"
        request := newPostWinRequest(player_name)
        response := httptest.NewRecorder()

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusInternalServerError)

        if len(store.WinCalls) != 0 {
            t.Errorf("got %d calls to RecordWin want %d", len(store.WinCalls), 0)
        }
    })

    t.Run("it records wins when POST", func(t *testing.T) {
        player_name := "Max"
        request := newPostWinRequest(player_name)
        response := httptest.NewRecorder()

        store.returnErrorDuringStore = false

        server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusAccepted)

        if len(store.WinCalls) != 1 {
            t.Errorf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
        }

        if store.WinCalls[0] != player_name {
            t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], player_name)
        }
    })
}

func newPostWinRequest(name string) *http.Request {
    req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
    return req
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    store := store.InMemoryPlayerStore{Scores: make(map[string]int)}
    server := PlayerServer{&store}
    player := "Pepper"

    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

    response := httptest.NewRecorder()
    server.ServeHTTP(response, newGetScoreRequest(player))
    assertStatus(t, response.Code, http.StatusOK)

    assertResponseBody(t, response.Body.String(), "3")
}