function validateString(input) {
    const validCharacters = /^[<>{}\[\]]*$/;
    if (!validCharacters.test(input)) {
        return false; 
    }

    const matchingPattern = /<[^<>]*>|{[^{}]*}|\[[^\[\]]*\]/g;
    let processedString = input;
    let matchFound;

    do {
        matchFound = processedString.match(matchingPattern);
        if (matchFound) {
            for (const match of matchFound) {
                processedString = processedString.replace(match, ''); // Menghapus kecocokan
            }
        }
    } while (matchFound);

    return processedString.length === 0;
}