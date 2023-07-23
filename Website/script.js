function getNum() {

    // Call the Go API
    fetch('http://localhost:8080/getNums')
        .then(response => response.json())
        .then(data => displayResults(data))
        .catch(error => console.error('Error:', error));

}

function displayResults(results) {
    const cont = document.getElementById('results');
    cont.innerHTML = '';

    results.forEach(result => {
        const resultDiv = document.createElement('div');
        bookDiv.innerHTML = `<strong>Number of occurencers of ${result.Word}:</strong> ${result.Number}`;
        cont.appendChild(resultDiv);
    });
}