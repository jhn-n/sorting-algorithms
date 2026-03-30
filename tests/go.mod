module sorting/tests

go 1.26.1

replace sorting/insertion => ../insertion

replace sorting/selection => ../selection

replace sorting/merge => ../merge

require (
	sorting/counting v0.0.0-00010101000000-000000000000
	sorting/insertion v0.0.0-00010101000000-000000000000
	sorting/merge v0.0.0-00010101000000-000000000000
	sorting/quick v0.0.0-00010101000000-000000000000
	sorting/selection v0.0.0-00010101000000-000000000000
)

replace sorting/counting => ../counting

replace sorting/quick => ../quick
