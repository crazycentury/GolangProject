
function findMatchingStrings(N, strings) {
    for (let i = 0; i < N; i++) {
        strings[i] = strings[i].toLowerCase();
    }

    for (let i = 0; i < N; i++) {
        for (let j = i + 1; j < N; j++) {
            if (strings[i] === strings[j]) {
                return `${i + 1} ${j + 1}`;
            }
        }
    }

    return false;
}