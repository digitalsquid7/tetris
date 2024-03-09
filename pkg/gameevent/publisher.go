package gameevent

type Subscriber interface {
	Update(events Events) error
}

type Publisher struct {
	subscribers []Subscriber
	events      Events
}

func NewPublisher() *Publisher {
	return &Publisher{
		events: MakeEvents(),
	}
}

func (p *Publisher) Subscribe(subscribers ...Subscriber) {
	for i := range subscribers {
		p.subscribers = append(p.subscribers, subscribers[i])
	}
}

func (p *Publisher) Publish(event Name) {
	p.events.Add(event)
}

func (p *Publisher) Notify() error {
	for i := range p.subscribers {
		err := p.subscribers[i].Update(p.events)
		if err != nil {
			return err
		}
	}

	p.events = MakeEvents()
	return nil
}
