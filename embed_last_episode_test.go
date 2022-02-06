package goembed

/* 
perlu diketahui hasil embed adalah permanent dan data file yg dibaca disimpan dalam binary filenya golang.
artinya bukan dilakukan secara realtime membaca file yg ada diluar.
hal ini menjadikan file binary golang sudah dicompile. kita tidak butuh lagi file externalnya,
dan bahkan jika diubah file externalnya, isi variablenya tidak akan berubah lagi.
 */

 