package extypes

import "github.com/tehsphinx/exalysis/gtpl"

//NewResponse creates a new response
func NewResponse() *Response {
	return &Response{}
}

//Response implements a response object. A response object is returned by the
//suggestor after examining a students solutions. It can produce a complete
//answer string to be replied to the student on exercism.
type Response struct {
	// the following slices make up the answer in the order given here
	greeting    []gtpl.Template
	intro       []gtpl.Template
	todo        []gtpl.Template
	improvement []gtpl.Template
	outro       []gtpl.Template
}

//SetGreeting sets the greeting overwriting already set or added greetings
func (s *Response) SetGreeting(template gtpl.Template) {
	s.greeting = []gtpl.Template{template}
}

//AppendGreeting adds a greeting template
func (s *Response) AppendGreeting(template gtpl.Template) {
	s.greeting = append(s.greeting, template)
}

//AppendIntro adds an intro template
func (s *Response) AppendIntro(template gtpl.Template) {
	s.intro = append(s.intro, template)
}

//AppendTodo adds a task to the list to be done before approval
func (s *Response) AppendTodo(template gtpl.Template) {
	s.todo = append(s.todo, template)
}

//AppendImprovement adds a optional improvement to the list that can be made
//by the student to improve the solution.
func (s *Response) AppendImprovement(template gtpl.Template) {
	s.improvement = append(s.improvement, template)
}

//AppendOutro adds an outro template
func (s *Response) AppendOutro(template gtpl.Template) {
	s.outro = append(s.outro, template)
}

//GetAnswerString returns the answer as a string to be used on exercism
func (s *Response) GetAnswerString() string {
	var answ string
	for _, t := range s.greeting {
		answ += t.TplString()
	}
	answ += s.praise().TplString()
	for _, t := range s.intro {
		answ += t.TplString()
	}

	var suggsAdded bool
	if len(s.todo) != 0 {
		answ += s.todoIntro().TplString()
		for _, t := range s.todo {
			answ += t.TplString()
		}
		suggsAdded = true
	}
	if len(s.improvement) != 0 {
		answ += s.improvementIntro().TplString()
		for _, t := range s.improvement {
			answ += t.TplString()
		}
		suggsAdded = true
	}
	if suggsAdded {
		s.AppendOutro(gtpl.Questions)
	}

	for _, t := range s.outro {
		answ += t.TplString()
	}
	return answ
}

func (s *Response) praise() gtpl.Template {
	var (
		l   = len(s.todo)*2 + len(s.improvement)
		adj string
	)
	switch {
	case l == 0:
		adj = "perfect"
	case l < 3:
		adj = "very good"
	case l < 6:
		adj = "good"
	default:
		adj = "interesting"
	}
	return gtpl.Praise.Format(adj)
}

func (s *Response) todoIntro() gtpl.Template {
	adj := "point"
	if 1 < len(s.todo) {
		adj = "points"
	}
	return gtpl.Todo.Format(adj)
}

func (s *Response) improvementIntro() gtpl.Template {
	adj := "one thought"
	if 1 < len(s.todo) {
		adj = "some thoughts"
	}
	return gtpl.Improvement.Format(adj)
}

//LenSuggestions returns the amount of suggestions added
func (s *Response) LenSuggestions() int {
	return len(s.todo) + len(s.improvement)
}

//GetTemplate returns requested template by id. Mainly used for testing to check if a template was
//being added or not on a specific example.
func (s *Response) GetTemplate(id string) (gtpl.Template, bool) {
	for _, t := range s.greeting {
		if t.ID() == id {
			return t, true
		}
	}
	for _, t := range s.intro {
		if t.ID() == id {
			return t, true
		}
	}
	if t, ok := s.GetSuggestion(id); ok {
		return t, ok
	}
	for _, t := range s.outro {
		if t.ID() == id {
			return t, true
		}
	}
	return nil, false
}

//GetSuggestion does the same as GetTemplate but only searches in todos and improments
func (s *Response) GetSuggestion(id string) (gtpl.Template, bool) {
	for _, t := range s.todo {
		if t.ID() == id {
			return t, true
		}
	}
	for _, t := range s.improvement {
		if t.ID() == id {
			return t, true
		}
	}
	return nil, false
}

//HasTemplate uses GetTemplate to search for a template but only returns if it was found or not
func (s *Response) HasTemplate(id string) bool {
	_, ok := s.GetTemplate(id)
	return ok
}

//HasSuggestion uses GetSuggestion to search for a template but only returns if it was found or not
func (s *Response) HasSuggestion(id string) bool {
	_, ok := s.GetTemplate(id)
	return ok
}
