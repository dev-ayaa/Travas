package controller

/*
func (tr *Travas) CreateTour() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tour model.Tour
		if err := ctx.Request.ParseForm(); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}
		tour.OperatorID = ctx.Request.Form.Get("operator_id")
		tour.TourTitle = ctx.Request.Form.Get("tour_title")
		tour.MeetingPoint = ctx.Request.Form.Get("meeting_point")
		tour.StartTime = ctx.Request.Form.Get("start_time")
		tour.LanguageOffered = ctx.Request.Form.Get("language_offered")
		tour.NumberOfTourist = ctx.Request.Form.Get("number_of_tourist")
		tour.Description = ctx.Request.Form.Get("description")
		tour.TourGuide = ctx.Request.Form.Get("tour_guide")
		tour.TourOperator = ctx.Request.Form.Get("tour_operator")
		tour.OperatorContact = ctx.Request.Form.Get("operator_contact")
		tour.Date = ctx.Request.Form.Get("date")

		tourID, err := tr.DB.InsertTour(tour)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("error while adding new user"))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"CreatedTour_ID": tourID,
			"data":           tour,
		})

	}
}

func (tr *Travas) DeleteTour() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		_, err := tr.DB.DeleteTour(id)
		if err != nil {
			ctx.JSON(406, gin.H{"message": "Tour could not be deleted", "error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(200, gin.H{"message": "Tour deleted"})
	}

}

func (tr *Travas) GetTour() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		tour, err := tr.DB.GetTour(id)
		if err != nil {
			ctx.JSON(404, gin.H{"message": "tour not found", "error": err.Error()})
			ctx.Abort()
		} else {
			ctx.JSON(200, gin.H{"data": tour})
		}
	}

}

func (tr *Travas) UpdateTour() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		tour := model.Tour{}

		if ctx.BindJSON(&tour) != nil {
			ctx.JSON(406, gin.H{"message": "Invalid Parameters"})
			ctx.Abort()
			return
		}
		_, err := tr.DB.UpdateTour(id, tour)
		if err != nil {
			ctx.JSON(406, gin.H{"message": "tour count not be updated", "error": err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(200, gin.H{"message": "tour updated"})

	}

}
func (tr *Travas) GetAllTours() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		list, err := tr.DB.FindAllTours()
		if err != nil {
			ctx.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
			ctx.Abort()
		} else {
			ctx.JSON(200, gin.H{"data": list})
		}
	}

}
*/
