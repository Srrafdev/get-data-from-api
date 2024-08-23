package box

func FilterLocaton(dataArtist []Data_Execute, loca string)*[]Data_Execute{
	var filter []Data_Execute
	for _, artist := range dataArtist {
		for _, valLoca := range artist.Locations{
			if valLoca == loca{
				filter = append(filter, artist)
			}
		}
		
	}
	return &filter
}
