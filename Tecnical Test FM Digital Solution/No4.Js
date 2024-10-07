function requestCuti(TotalCutiBersama, firstDay, tanggalCuti, durasiCuti) {

    const CUTI_KANTOR = 14;
    const CUTI_MAKSIMAL_BERTURUTAN = 3;

    const joinDate = new Date(firstDay);
    const cutiDate = new Date(tanggalCuti);

    const tanggalAwalCutiPribadi = new Date(joinDate);
    tanggalAwalCutiPribadi.setDate(tanggalAwalCutiPribadi.getDate() + 180);

    if (cutiDate < tanggalAwalCutiPribadi) {
        return [false, "Karena belum 180 hari sejak tanggal join karyawan."];
    }

    const akhirTahun = new Date(cutiDate.getFullYear(), 11, 31); // 31 Desember
    const jumlahHari = Math.floor((akhirTahun - tanggalAwalCutiPribadi) / (1000 * 60 * 60 * 24)) + 1; // +1 untuk menghitung hari terakhir

    const jumlahCutiPribadi = Math.floor((jumlahHari / 365) * (CUTI_KANTOR - TotalCutiBersama));

    if (durasiCuti > CUTI_MAKSIMAL_BERTURUTAN) {
        return [false, `Durasi cuti tidak boleh lebih dari ${CUTI_MAKSIMAL_BERTURUTAN} hari berturutan.`];
    }

    if (durasiCuti > jumlahCutiPribadi) {
        return [false, "Jumlah cuti pribadi tidak mencukupi."];
    }

    return [true, "Karyawan dapat mengambil cuti pribadi."];
}


const [result1, alasan1] = requestCuti(7, "2021-05-01", "2021-07-05", 1);
console.log(result1, alasan1);

const [result2, alasan2] = requestCuti(7, "2021-01-05", "2021-12-18", 1);
console.log(result2, alasan2);
