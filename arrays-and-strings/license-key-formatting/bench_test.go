package license_key_formatting

import "testing"

func Benchmark_Sol(b *testing.B) {
	input := "5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh"
	for i := 0; i < b.N; i++ {
		licenseKeyFormatting(input, 4)
	}
}

func Benchmark_Fast(b *testing.B) {
	input := "5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh-5F3Z-2e-9-w-asdsad-dsadasdsad--dadasdsadsa-dsad-asd-sad-sad-as-fesg-tj-ghfk-h-gikx-dg-yaer-sy-srdh-dth-xdh-xfgh-xfgh-xfgh"
	for i := 0; i < b.N; i++ {
		licenseKeyFormattingNoFast(input, 4)
	}
}
