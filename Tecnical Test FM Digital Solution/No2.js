function calculateChange(totalBelanja, uangDibayar) {
    const pecahanUang = [100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100];

    if (uangDibayar < totalBelanja) {
        return 'false, kurang bayar'; 
    }

    let kembalian = uangDibayar - totalBelanja;
    kembalian = Math.floor(kembalian / 100) * 100;

    let hasilPecahan = {};

    for (let i = 0; i < pecahanUang.length; i++) {
        if (kembalian === 0) break; // Jika tidak ada kembalian lagi
        
        const jumlahUang = Math.floor(kembalian / pecahanUang[i]); // Jumlah pecahan yang diberikan
        if (jumlahUang > 0) {
            hasilPecahan[pecahanUang[i]] = jumlahUang; // Simpan hasil pecahan
            kembalian -= jumlahUang * pecahanUang[i]; // Kurangi kembalian
        }
    }

    
    let output = `Kembalian yang harus diberikan kasir: ${uangDibayar - totalBelanja},\ndibulatkan menjadi ${kembalian}\nPecahan uang:\n`;
    for (const [pecahan, jumlah] of Object.entries(hasilPecahan)) {
        output += `${jumlah} lembar ${pecahan}\n`;
    }

    return output.trim();
}