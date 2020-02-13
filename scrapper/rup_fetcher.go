package scrapper

// func (f *collyScrapper) Save(ctx context.Context, rups []model.RupItem) error {
// 	return f.svc.SaveRup(ctx, rups)
// }

// func (f *collyScrapper) Fetch(ctx context.Context, opt model.RupOptions) ([]model.RupItem, error) {
// 	res, err := f.fetchRup(opt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	rups := make([]model.RupItem, 0)

// 	for _, item := range res.AaData {
// 		/*	response Penyedia
// 							[
// 					0     "19405324",
// 					1     "Belanja Bahan Perbekalan Habis Pakai",
// 					2     "500000000",
// 					3     "E-Purchasing",
// 					4     "APBD",
// 					5     "19405324",
// 					6     "February 2019"
// 					   ],
// 					response swakelola
// 				[
// 				    0  "19889649",
// 				    1  "Peningkatan Pelayanan dan Kinerja Pemda",
// 				    2  "Peningkatan Disiplin Aparatur",
// 				    3  "159987000",
// 				    4  "APBD",
// 				    5  "19889649",
// 				    6  "January 2019"
// 				    ],
// 					response penyedia dalam swakelola
// 			[
// 			    0  "19137689",
// 			    1  "Belanja Makanan dan Minuman",
// 			    2  "73050000",
// 			    3  "Pengadaan Langsung",
// 			    4  "APBD, APBD",
// 			    5  "19137689",
// 			    6  "Januari 2019",
// 			    7  "true",
// 			    8  "true",
// 			    9  "19137689"
// 			    ],
// 		*/

// 		rup := model.RupItem{}
// 		rup.KodeOpd = opt.KodeOpd
// 		rup.Tahun = opt.Tahun
// 		rup.Kategori = opt.Kategori
// 		//response
// 		rup.KodeRup = item[0]
// 		if opt.Kategori == model.KategoriSwakelola {
// 			rup.Kegiatan = &item[1]
// 			rup.NamaPaket = item[2]
// 			rup.Pagu = item[3]
// 			if opt.Metode == nil {
// 				m := model.MetodeSwakelola
// 				rup.Metode = m
// 			} else {
// 				rup.Metode = *opt.Metode
// 			}

// 		} else {
// 			rup.NamaPaket = item[1]
// 			rup.Pagu = item[2]
// 			metode := model.Metode(item[3])
// 			// rup.Metode = item[3]
// 			rup.Metode = metode
// 		}

// 		rup.SumberDana = item[4]
// 		//rup.KodeRUP = item[5] //double result
// 		rup.Waktu = item[6]
// 		rups = append(rups, rup)
// 	}
// 	return rups, nil

// }
