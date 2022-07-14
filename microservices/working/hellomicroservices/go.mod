module animo.com/hellomicroservices

go 1.18

require animo.com/handlers v0.0.0-00010101000000-000000000000

require animo.com/data v0.0.0-00010101000000-000000000000 // indirect

//#replace animo/handlers => ../handlers

replace animo.com/handlers => ../handlers

replace animo.com/data => ../data
