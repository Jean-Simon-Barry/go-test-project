package book

// Service : service for books
type Service struct {
	BookRepository Repository
}

// ProvideBookService : provides ProvideProductService
func ProvideBookService(p Repository) Service {
	return Service{BookRepository: p}
}

// FindBookByID : finds book by id
func (s *Service) FindBookByID(bookID uint) Book {
	return s.BookRepository.FindBookByID(bookID)
}

// FindBookByTitle : finds book by title
func (s *Service) FindBookByTitle(bookTitle string) Book {
	return s.BookRepository.FindBookByTitle(bookTitle)
}

// GetAllBooks : gets all books from db (mainly for testing)
func (s *Service) GetAllBooks() []Book {
	return s.BookRepository.GetAllBooks()
}
