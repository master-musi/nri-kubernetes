package src

//TODO NamespaceFilterer is temporarilly here but it's just to not stop our development.
type NamespaceFilterer interface {
	IsAllowed(namespace string) bool
}

type NamespaceFilter struct{}

func (nf NamespaceFilter) IsAllowed(namespace string) bool {
	//TODO I've forced it to newrelic to check that it filters the newrelic namespace
	if namespace == "newrelic" {
		return false
	}
	return true
}
