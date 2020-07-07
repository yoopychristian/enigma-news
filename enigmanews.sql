create database enigmanews;

create table registrasi(
id varchar(36) primary key,
email varchar(40),
full_name varchar (50),
phone_number int,
username varchar(40),
password varchar(50));

create table user(
id varchar(36),
username varchar(40),
password varchar(50));

create table article(
id varchar(36),
title text(50),
article longtext);

select * from registrasi;
select * from article;

insert into registrasi values
('B001', 'alexaqwe@gmail.com', 'alex nurdin', '08112391299', 'alexcuy', 'pwpwpw'),
('B002', 'doniara@gmail.com', 'doni aradong', '08129381929', 'doniara', 'suesue');

insert into article values
('ART001', 'Bahaya Mie Instan Bagi Kesehatan', 'Setiap orang sudah pasti kenal dengan makanan cepat saji bernama mie instan Kecuali mudah diperoleh, variasi makanan yang satu ini juga mempunyai rasa yang sedap.Hanya saja, dibalik sifat praktis dan rasanya yang sangat sedap, mie instan juga mempunyai sejumlah resiko.

Menurut sejumlah hasil penelitian, Terlalu kerap mengonsumsi mie instan dapat meningkatkan resiko timbulnya penyakit kanker, ginjal dan usus buntu. Pada beberapa kasus orang yang mengkonsumsi mie instan berlebihan juga dapat menyebabkan kegemukan atau obesitas.

Sekiranya Anda umumnya mengonsumsi mie instan tiap-tiap hari, mulailah untuk menguranginya dengan tenggang waktu 2hingga 3 hari dan lakukan sampai Anda terbebas dari mie instan sama sekali. Menurut penelitian rupanya didalam mie instan terdapat kandungan lilin yang sangat membahayakan bagi kesehatan manusia.

Mie instan mengandung kandungan lilin yang berguna untuk membuat mie tak lengket satu dengan lainnya. Dengan seringnya kita mengonsumsi mie instan berarti kita juga sudah memasukkan banyak kandungan lilin ke dalam tubuh kita. Dengan seringnya kita mengonsumsi mie instan berarti kita juga sudah memasukkan banyak kandungan lilin ke dalam tubuh kita. Kandungan lilin tersebut akan merusak metode kerja pencernaan dalam tubuh karena baru dapat dicerna oleh tubuh dalam waktu minimal 2 hari.');

insert into article values
('ART002', 'PENDIDIKAN SANGAT BERPENGARUH TERHADAP
PENINGKATAN KUALITAS HIDUP SUKU ANAK DALAM', 'Sebuah tulisan sangat penting bagi setiap makhluk hidup di dunia salah satunya yaitu kita sebagai manusia. Karena manusia sebagai makhluk hidup yang memiliki nalar dan pikiran untuk meningkatkan kesejahteraan manusia itu sendiri. Tulisan dapat mensupport kehidupan seseorang maupun kategori orang.

Di dalam sebuah pemikiran secara universal suatu pendidikan dan pengajaran terhadap setiap manusia sangat penting tidak terkecuali bagi suku anak dalam. Suku anak dalam sangatlah membutuhkan pengajaran yang cocok untuk dapat meningkatkan kesejahteraannya dalam menjalani kehidupan.

Terkait pembelajaran tentang analisis data yang dilakukan secara literatur ditemukan sebagai persoalan mengenai minimnya pengajaran di tingkat suku anak dalam. Dari hasil analisis tersebut di harapkan supaya pemerintah dapat menciptakan pengajaran yang mengarah terhadap pengaplikasian pengajaran di daerah pelosok suku anak dalam.'),
('ART003', 'Karya Tentang Pendidikan', 'Sebagai contoh pada era modernisasi sekarang ini para suku anak dalam seperti terpinggirkan. Dalam hal pengajaran tak hanya dipikirkan oleh pemerintah saja yang harus bergerak namun sesama manusia juga semestinya saling menolong menyangkut kesejahteraan bersama. Dalam kehidupan sehari-hari kita perlu mengetahui pengetahuan tentang tata tertib dari dewa yang menguasai alam ikut memberi pengaruh pola hidup Orang Rimba, terpenting dalam mengelola alam sekitar.

Orang Rimba sangat menghargai dan terikat dengan lingkungan sekitar (hutan), dimana mereka makan juga minum dari apa yang disediakan di hutan. Bagi Orang Rimba, hutan yaitu komponen dari hidup mereka yang semestinya di lindungi serta orang rimba juga mempunyai motto “hutan yaitu kehidupan dan kehidupan yaitu hutan”. Keduanya berjalan seiring dan mereka tak pernah mengharapkan untuk hidup diluar hutan karena hutan dirasakan sudah cukup memenuhi kebutuhan hidup mereka (Lucky Ayu Wulandari, 2009).

Selain itu dilihat dari kehidupan mereka yang sangat bertumpu pada alam mereka juga semestinya di berikan pengajaran yang cocok. Untuk menerima kehidupan yang sejahtera, mereka juga memakai pengajaran mereka sanggup menjaga keseimbangan ekosistem di alam. Dan juga mereka dapat memenuhi kebutuhan sehari-hari yang sangat bertumpu pada ekosistem alam. Dimana alam harus dijaga supaya tidak punah dan dapat hidup selayaknya masyarakat modern zaman sekarang.');
