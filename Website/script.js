function getNum() {
    const val = document.querySelector('input').value;

    // Call the Go API
    fetch('http://localhost:8080/getNums', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: `{
           word: word 
        }`,})
        .then(response => response.json())
        .then(data => displayResults(data))
        .catch(error => console.error('Error:', error));

}

function displayResults(results) {
    // Get the element to write the results to
    const cont = document.getElementById('results');

    cont.innerHTML = '';

    // For each result write the result to the element
    results.forEach(result => {
        const resultDiv = document.createElement('div');
        bookDiv.innerHTML = `<strong>Number of occurencers of ${result.Word}:</strong> ${result.Number}`;
        cont.appendChild(resultDiv);
    });
}