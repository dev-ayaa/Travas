package repo

//All this will be use in the tour operator website
/*
func (td *TravasDB) InsertTour(tour model.Tour) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var id primitive.ObjectID

	filter := bson.D{{Key: "tour_title", Value: tour.TourTitle}}
	var res bson.M
	err := TouristData(td.DB, "tours").FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			tour.ID = primitive.NewObjectID()
			doc := bson.D{
				{Key: "_id", Value: tour.ID},
				{Key: "operator_id", Value: tour.OperatorID},
				{Key: "tour_title", Value: tour.TourTitle},
				{Key: "meeeting_point", Value: tour.MeetingPoint},
				{Key: "start_time", Value: tour.StartTime},
				{Key: "language_offered", Value: tour.LanguageOffered},
				{Key: "number_of_tourist", Value: tour.NumberOfTourist},
				{Key: "description", Value: tour.Description},
				{Key: "tour_guide", Value: tour.TourGuide},
				{Key: "tour_operator", Value: tour.TourOperator},
				{Key: "operator_contact", Value: tour.OperatorContact},
				{Key: "date", Value: tour.Date},
			}
			_, insertErr := TouristData(td.DB, "tours").InsertOne(ctx, doc)
			if insertErr != nil {
				td.App.ErrorLogger.Fatalf("cannot add user to the database : %v ", insertErr)
			}
			return tour.ID, nil
		}
		td.App.ErrorLogger.Fatal(err)
	}

	for k, v := range res {
		if k == "_id" {
			switch tourID := v.(type) {
			case primitive.ObjectID:
				id = tourID
			}
		}
	}
	return id, err
}

func (td *TravasDB) DeleteTour(tourID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result bson.M

	filter := bson.D{{Key: "_id", Value: tourID}}
	err := TouristData(td.DB, "tours").FindOneAndDelete(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			td.App.ErrorLogger.Println("no document found for this query")
			return false, err
		}
		td.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}

	return true, nil
}

func (td *TravasDB) UpdateTour(tourID string, tour model.Tour) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: tourID}}
	_, err := TouristData(td.DB, "tours").UpdateByID(ctx, filter, tour)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (td *TravasDB) GetTour(tourID string) (tour model.Tour, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: tourID}}
	err = TouristData(td.DB, "tours").FindOne(ctx, filter).Decode(&tour)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			td.App.ErrorLogger.Println("no document found for this query")
		}
		td.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}
	return tour, err
}
func (td *TravasDB) FindAllTours() (tours []model.Tour, err error) {

	cur, err := TouristData(td.DB, "tours").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	if err = cur.All(context.Background(), &tours); err != nil {
		log.Fatal(err)
	}

	return tours, err

}
*/
