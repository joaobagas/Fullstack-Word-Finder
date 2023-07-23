function getNum() {

    // Call the Go API
    fetch('http://localhost:8080/books')
        .then(response => response.json())
        .then(data => displayBooks(data))
        .catch(error => console.error('Error:', error));

}

function displayBooks(results) {
    const booksContainer = document.getElementById('results');
    booksContainer.innerHTML = '';

    results.forEach(result => {
        const bookDiv = document.createElement('div');

        // Here you create the line with the results!
        bookDiv.innerHTML = `<strong>Number of occurencers of ${result.word}:</strong> ${result.number}`;

        booksContainer.appendChild(bookDiv);
    });
}