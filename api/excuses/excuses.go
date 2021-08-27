package excuses

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Excuses struct {
	list []string
}

type RandomResponse struct {
	Excuse string
}

type CountResponse struct {
	Count int
}

type LoadRequest struct {
	URL string
}

func New() *Excuses {
	return &Excuses{}
}

func (e *Excuses) Default(ctx context.Context) (*RandomResponse, error) {
	return e.Random(ctx)
}

// Count returns number of available excuses.
func (e *Excuses) Count(ctx context.Context) (*CountResponse, error) {
	return &CountResponse{Count: len(e.list)}, nil
}

// Clear removes all excuses
func (e *Excuses) Clear(ctx context.Context) {
	e.list = make([]string, 0)
}

// Random returns one excuse from the list.
// Or error if list is empty.
func (e *Excuses) Random(ctx context.Context) (*RandomResponse, error) {
	l := len(e.list)
	if l == 0 {
		return nil, fmt.Errorf("no excuses")
	}
	return &RandomResponse{Excuse: e.list[rand.Intn(l)]}, nil
}

// Load adds new excuses to the list.
// It loads list of strings from the provided URL
// checks if we already have that if not adds it.
func (e *Excuses) Load(ctx context.Context, req LoadRequest) (*CountResponse, error) {
	if req.URL == "" {
		return nil, fmt.Errorf("URL not found")
	}

	rsp, err := http.Get(req.URL)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	log.Printf("count before: %d", len(e.list))
	scanner := bufio.NewScanner(rsp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if e.exists(line) {
			log.Printf("this one already exists: %s", line)
			continue
		}
		e.list = append(e.list, line)
		log.Printf("added: %s", line)
	}
	log.Printf("count after: %d", len(e.list))

	return e.Count(ctx)
}

// exists checks if we already have that excuse
func (e *Excuses) exists(line string) bool {
	for _, v := range e.list {
		if v == line {
			return true
		}
	}
	return false
}
