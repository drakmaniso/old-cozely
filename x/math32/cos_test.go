// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package math32

import (
	"math"
	"testing"
)

//------------------------------------------------------------------------------

func TestCos(t *testing.T) {
	var x float32
	for _, tt := range cosTests {
		x = Cos(tt.in)
		if !IsNearlyEqual(x, tt.out, EpsilonFloat32) {
			t.Errorf("Relative error for Cos(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
		if !IsAlmostEqual(x, tt.out, 2) {
			t.Errorf("ULP error for Cos(%.100g): %.100g instead of %.100g\n", tt.in, x, tt.out)
		}
	}
}

//------------------------------------------------------------------------------

func BenchmarkCos_math(b *testing.B) {
	a := float64(0.5)
	for i := 0; i < b.N; i++ {
		result64 = math.Cos(a)
	}
}

//------------------------------------------------------------------------------

func BenchmarkCos_float32math(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		result = float32(math.Cos(float64(a)))
	}
}

//------------------------------------------------------------------------------

func BenchmarkCos_glam(b *testing.B) {
	a := float32(0.5)
	for i := 0; i < b.N; i++ {
		result = Cos(a)
	}
}

//------------------------------------------------------------------------------

var cosTests = [...]struct {
	in  float32
	out float32
}{
	// Special cases

	{0, 1},
	{Pi / 6, 0.8660254037844386467637231707529},
	{Pi / 4, 0.7071067811865475244008443621048},
	{Pi / 3, 0.5},
	{Pi / 2, -4.37113883e-8},

	// The following values have been generated with Gnu MPFR, 24 bits mantissa, roundTiesToEven

	{-3.14000010e+00, -9.99998748e-01},
	{-3.13000011e+00, -9.99932826e-01},
	{-3.12000012e+00, -9.99766886e-01},
	{-3.11000013e+00, -9.99500990e-01},
	{-3.10000014e+00, -9.99135137e-01},
	{-3.09000015e+00, -9.98669386e-01},
	{-3.08000016e+00, -9.98103797e-01},
	{-3.07000017e+00, -9.97438371e-01},
	{-3.06000018e+00, -9.96673167e-01},
	{-3.05000019e+00, -9.95808363e-01},
	{-3.04000020e+00, -9.94843900e-01},
	{-3.03000021e+00, -9.93780017e-01},
	{-3.02000022e+00, -9.92616773e-01},
	{-3.01000023e+00, -9.91354227e-01},
	{-3.00000024e+00, -9.89992559e-01},
	{-2.99000025e+00, -9.88531888e-01},
	{-2.98000026e+00, -9.86972332e-01},
	{-2.97000027e+00, -9.85314131e-01},
	{-2.96000028e+00, -9.83557343e-01},
	{-2.95000029e+00, -9.81702268e-01},
	{-2.94000030e+00, -9.79748964e-01},
	{-2.93000031e+00, -9.77697730e-01},
	{-2.92000031e+00, -9.75548744e-01},
	{-2.91000032e+00, -9.73302126e-01},
	{-2.90000033e+00, -9.70958233e-01},
	{-2.89000034e+00, -9.68517244e-01},
	{-2.88000035e+00, -9.65979397e-01},
	{-2.87000036e+00, -9.63344932e-01},
	{-2.86000037e+00, -9.60614204e-01},
	{-2.85000038e+00, -9.57787335e-01},
	{-2.84000039e+00, -9.54864740e-01},
	{-2.83000040e+00, -9.51846659e-01},
	{-2.82000041e+00, -9.48733330e-01},
	{-2.81000042e+00, -9.45525169e-01},
	{-2.80000043e+00, -9.42222476e-01},
	{-2.79000044e+00, -9.38825548e-01},
	{-2.78000045e+00, -9.35334742e-01},
	{-2.77000046e+00, -9.31750417e-01},
	{-2.76000047e+00, -9.28072870e-01},
	{-2.75000048e+00, -9.24302578e-01},
	{-2.74000049e+00, -9.20439780e-01},
	{-2.73000050e+00, -9.16485012e-01},
	{-2.72000051e+00, -9.12438571e-01},
	{-2.71000051e+00, -9.08300877e-01},
	{-2.70000052e+00, -9.04072344e-01},
	{-2.69000053e+00, -8.99753451e-01},
	{-2.68000054e+00, -8.95344555e-01},
	{-2.67000055e+00, -8.90846133e-01},
	{-2.66000056e+00, -8.86258602e-01},
	{-2.65000057e+00, -8.81582439e-01},
	{-2.64000058e+00, -8.76818180e-01},
	{-2.63000059e+00, -8.71966183e-01},
	{-2.62000060e+00, -8.67027044e-01},
	{-2.61000061e+00, -8.62001121e-01},
	{-2.60000062e+00, -8.56889069e-01},
	{-2.59000063e+00, -8.51691306e-01},
	{-2.58000064e+00, -8.46408367e-01},
	{-2.57000065e+00, -8.41040790e-01},
	{-2.56000066e+00, -8.35589111e-01},
	{-2.55000067e+00, -8.30053926e-01},
	{-2.54000068e+00, -8.24435651e-01},
	{-2.53000069e+00, -8.18735003e-01},
	{-2.52000070e+00, -8.12952459e-01},
	{-2.51000071e+00, -8.07088614e-01},
	{-2.50000072e+00, -8.01144063e-01},
	{-2.49000072e+00, -7.95119405e-01},
	{-2.48000073e+00, -7.89015174e-01},
	{-2.47000074e+00, -7.82832146e-01},
	{-2.46000075e+00, -7.76570737e-01},
	{-2.45000076e+00, -7.70231724e-01},
	{-2.44000077e+00, -7.63815701e-01},
	{-2.43000078e+00, -7.57323265e-01},
	{-2.42000079e+00, -7.50755131e-01},
	{-2.41000080e+00, -7.44111896e-01},
	{-2.40000081e+00, -7.37394273e-01},
	{-2.39000082e+00, -7.30602860e-01},
	{-2.38000083e+00, -7.23738432e-01},
	{-2.37000084e+00, -7.16801643e-01},
	{-2.36000085e+00, -7.09793150e-01},
	{-2.35000086e+00, -7.02713668e-01},
	{-2.34000087e+00, -6.95563972e-01},
	{-2.33000088e+00, -6.88344657e-01},
	{-2.32000089e+00, -6.81056559e-01},
	{-2.31000090e+00, -6.73700273e-01},
	{-2.30000091e+00, -6.66276693e-01},
	{-2.29000092e+00, -6.58786476e-01},
	{-2.28000093e+00, -6.51230335e-01},
	{-2.27000093e+00, -6.43609107e-01},
	{-2.26000094e+00, -6.35923564e-01},
	{-2.25000095e+00, -6.28174365e-01},
	{-2.24000096e+00, -6.20362341e-01},
	{-2.23000097e+00, -6.12488329e-01},
	{-2.22000098e+00, -6.04553044e-01},
	{-2.21000099e+00, -5.96557319e-01},
	{-2.20000100e+00, -5.88501930e-01},
	{-2.19000101e+00, -5.80387712e-01},
	{-2.18000102e+00, -5.72215378e-01},
	{-2.17000103e+00, -5.63985884e-01},
	{-2.16000104e+00, -5.55700004e-01},
	{-2.15000105e+00, -5.47358572e-01},
	{-2.14000106e+00, -5.38962364e-01},
	{-2.13000107e+00, -5.30512214e-01},
	{-2.12000108e+00, -5.22009075e-01},
	{-2.11000109e+00, -5.13453722e-01},
	{-2.10000110e+00, -5.04847050e-01},
	{-2.09000111e+00, -4.96189862e-01},
	{-2.08000112e+00, -4.87483084e-01},
	{-2.07000113e+00, -4.78727520e-01},
	{-2.06000113e+00, -4.69924122e-01},
	{-2.05000114e+00, -4.61073697e-01},
	{-2.04000115e+00, -4.52177197e-01},
	{-2.03000116e+00, -4.43235457e-01},
	{-2.02000117e+00, -4.34249401e-01},
	{-2.01000118e+00, -4.25219923e-01},
	{-2.00000119e+00, -4.16147918e-01},
	{-1.99000120e+00, -4.07034308e-01},
	{-1.98000121e+00, -3.97879988e-01},
	{-1.97000122e+00, -3.88685882e-01},
	{-1.96000123e+00, -3.79452914e-01},
	{-1.95000124e+00, -3.70181978e-01},
	{-1.94000125e+00, -3.60874057e-01},
	{-1.93000126e+00, -3.51530015e-01},
	{-1.92000127e+00, -3.42150837e-01},
	{-1.91000128e+00, -3.32737446e-01},
	{-1.90000129e+00, -3.23290795e-01},
	{-1.89000130e+00, -3.13811779e-01},
	{-1.88000131e+00, -3.04301411e-01},
	{-1.87000132e+00, -2.94760615e-01},
	{-1.86000133e+00, -2.85190344e-01},
	{-1.85000134e+00, -2.75591522e-01},
	{-1.84000134e+00, -2.65965164e-01},
	{-1.83000135e+00, -2.56312221e-01},
	{-1.82000136e+00, -2.46633634e-01},
	{-1.81000137e+00, -2.36930385e-01},
	{-1.80000138e+00, -2.27203444e-01},
	{-1.79000139e+00, -2.17453778e-01},
	{-1.78000140e+00, -2.07682371e-01},
	{-1.77000141e+00, -1.97890192e-01},
	{-1.76000142e+00, -1.88078240e-01},
	{-1.75000143e+00, -1.78247467e-01},
	{-1.74000144e+00, -1.68398872e-01},
	{-1.73000145e+00, -1.58533439e-01},
	{-1.72000146e+00, -1.48652136e-01},
	{-1.71000147e+00, -1.38755992e-01},
	{-1.70000148e+00, -1.28845960e-01},
	{-1.69000149e+00, -1.18923046e-01},
	{-1.68000150e+00, -1.08988240e-01},
	{-1.67000151e+00, -9.90425348e-02},
	{-1.66000152e+00, -8.90869275e-02},
	{-1.65000153e+00, -7.91224092e-02},
	{-1.64000154e+00, -6.91499785e-02},
	{-1.63000154e+00, -5.91706373e-02},
	{-1.62000155e+00, -4.91853729e-02},
	{-1.61000156e+00, -3.91951948e-02},
	{-1.60000157e+00, -2.92010959e-02},
	{-1.59000158e+00, -1.92040764e-02},
	{-1.58000159e+00, -9.20513552e-03},
	{-1.57000160e+00, +7.94724561e-04},
	{-1.56000161e+00, +1.07945055e-02},
	{-1.55000162e+00, +2.07932070e-02},
	{-1.54000163e+00, +3.07898298e-02},
	{-1.53000164e+00, +4.07833718e-02},
	{-1.52000165e+00, +5.07728383e-02},
	{-1.51000166e+00, +6.07572235e-02},
	{-1.50000167e+00, +7.07355365e-02},
	{-1.49000168e+00, +8.07067752e-02},
	{-1.48000169e+00, +9.06699449e-02},
	{-1.47000170e+00, +1.00624047e-01},
	{-1.46000171e+00, +1.10568084e-01},
	{-1.45000172e+00, +1.20501064e-01},
	{-1.44000173e+00, +1.30421996e-01},
	{-1.43000174e+00, +1.40329883e-01},
	{-1.42000175e+00, +1.50223747e-01},
	{-1.41000175e+00, +1.60102576e-01},
	{-1.40000176e+00, +1.69965401e-01},
	{-1.39000177e+00, +1.79811239e-01},
	{-1.38000178e+00, +1.89639077e-01},
	{-1.37000179e+00, +1.99447960e-01},
	{-1.36000180e+00, +2.09236905e-01},
	{-1.35000181e+00, +2.19004914e-01},
	{-1.34000182e+00, +2.28751034e-01},
	{-1.33000183e+00, +2.38474280e-01},
	{-1.32000184e+00, +2.48173669e-01},
	{-1.31000185e+00, +2.57848233e-01},
	{-1.30000186e+00, +2.67497033e-01},
	{-1.29000187e+00, +2.77119070e-01},
	{-1.28000188e+00, +2.86713421e-01},
	{-1.27000189e+00, +2.96279073e-01},
	{-1.26000190e+00, +3.05815101e-01},
	{-1.25000191e+00, +3.15320551e-01},
	{-1.24000192e+00, +3.24794471e-01},
	{-1.23000193e+00, +3.34235907e-01},
	{-1.22000194e+00, +3.43643934e-01},
	{-1.21000195e+00, +3.53017569e-01},
	{-1.20000196e+00, +3.62355918e-01},
	{-1.19000196e+00, +3.71658057e-01},
	{-1.18000197e+00, +3.80923003e-01},
	{-1.17000198e+00, +3.90149862e-01},
	{-1.16000199e+00, +3.99337709e-01},
	{-1.15000200e+00, +4.08485621e-01},
	{-1.14000201e+00, +4.17592674e-01},
	{-1.13000202e+00, +4.26657975e-01},
	{-1.12000203e+00, +4.35680628e-01},
	{-1.11000204e+00, +4.44659680e-01},
	{-1.10000205e+00, +4.53594297e-01},
	{-1.09000206e+00, +4.62483555e-01},
	{-1.08000207e+00, +4.71326530e-01},
	{-1.07000208e+00, +4.80122417e-01},
	{-1.06000209e+00, +4.88870263e-01},
	{-1.05000210e+00, +4.97569233e-01},
	{-1.04000211e+00, +5.06218433e-01},
	{-1.03000212e+00, +5.14817059e-01},
	{-1.02000213e+00, +5.23364127e-01},
	{-1.01000214e+00, +5.31858921e-01},
	{-1.00000215e+00, +5.40300488e-01},
	{-9.90002155e-01, +5.48688054e-01},
	{-9.80002165e-01, +5.57020724e-01},
	{-9.70002174e-01, +5.65297723e-01},
	{-9.60002184e-01, +5.73518217e-01},
	{-9.50002193e-01, +5.81681311e-01},
	{-9.40002203e-01, +5.89786232e-01},
	{-9.30002213e-01, +5.97832203e-01},
	{-9.20002222e-01, +6.05818391e-01},
	{-9.10002232e-01, +6.13743961e-01},
	{-9.00002241e-01, +6.21608198e-01},
	{-8.90002251e-01, +6.29410267e-01},
	{-8.80002260e-01, +6.37149394e-01},
	{-8.70002270e-01, +6.44824803e-01},
	{-8.60002279e-01, +6.52435720e-01},
	{-8.50002289e-01, +6.59981430e-01},
	{-8.40002298e-01, +6.67461097e-01},
	{-8.30002308e-01, +6.74874067e-01},
	{-8.20002317e-01, +6.82219505e-01},
	{-8.10002327e-01, +6.89496756e-01},
	{-8.00002337e-01, +6.96705043e-01},
	{-7.90002346e-01, +7.03843653e-01},
	{-7.80002356e-01, +7.10911870e-01},
	{-7.70002365e-01, +7.17909038e-01},
	{-7.60002375e-01, +7.24834383e-01},
	{-7.50002384e-01, +7.31687248e-01},
	{-7.40002394e-01, +7.38466918e-01},
	{-7.30002403e-01, +7.45172799e-01},
	{-7.20002413e-01, +7.51804113e-01},
	{-7.10002422e-01, +7.58360326e-01},
	{-7.00002432e-01, +7.64840603e-01},
	{-6.90002441e-01, +7.71244466e-01},
	{-6.80002451e-01, +7.77571201e-01},
	{-6.70002460e-01, +7.83820152e-01},
	{-6.60002470e-01, +7.89990723e-01},
	{-6.50002480e-01, +7.96082318e-01},
	{-6.40002489e-01, +8.02094281e-01},
	{-6.30002499e-01, +8.08026016e-01},
	{-6.20002508e-01, +8.13876987e-01},
	{-6.10002518e-01, +8.19646597e-01},
	{-6.00002527e-01, +8.25334191e-01},
	{-5.90002537e-01, +8.30939293e-01},
	{-5.80002546e-01, +8.36461246e-01},
	{-5.70002556e-01, +8.41899574e-01},
	{-5.60002565e-01, +8.47253740e-01},
	{-5.50002575e-01, +8.52523148e-01},
	{-5.40002584e-01, +8.57707381e-01},
	{-5.30002594e-01, +8.62805784e-01},
	{-5.20002604e-01, +8.67817879e-01},
	{-5.10002613e-01, +8.72743249e-01},
	{-5.00002623e-01, +8.77581298e-01},
	{-4.90002632e-01, +8.82331610e-01},
	{-4.80002642e-01, +8.86993706e-01},
	{-4.70002651e-01, +8.91567111e-01},
	{-4.60002661e-01, +8.96051288e-01},
	{-4.50002670e-01, +9.00445938e-01},
	{-4.40002680e-01, +9.04750526e-01},
	{-4.30002689e-01, +9.08964634e-01},
	{-4.20002699e-01, +9.13087845e-01},
	{-4.10002708e-01, +9.17119741e-01},
	{-4.00002718e-01, +9.21059906e-01},
	{-3.90002728e-01, +9.24908042e-01},
	{-3.80002737e-01, +9.28663611e-01},
	{-3.70002747e-01, +9.32326376e-01},
	{-3.60002756e-01, +9.35895860e-01},
	{-3.50002766e-01, +9.39371765e-01},
	{-3.40002775e-01, +9.42753732e-01},
	{-3.30002785e-01, +9.46041465e-01},
	{-3.20002794e-01, +9.49234545e-01},
	{-3.10002804e-01, +9.52332735e-01},
	{-3.00002813e-01, +9.55335677e-01},
	{-2.90002823e-01, +9.58243072e-01},
	{-2.80002832e-01, +9.61054683e-01},
	{-2.70002842e-01, +9.63770151e-01},
	{-2.60002851e-01, +9.66389239e-01},
	{-2.50002861e-01, +9.68911707e-01},
	{-2.40002856e-01, +9.71337318e-01},
	{-2.30002850e-01, +9.73665774e-01},
	{-2.20002845e-01, +9.75896835e-01},
	{-2.10002840e-01, +9.78030324e-01},
	{-2.00002834e-01, +9.80066001e-01},
	{-1.90002829e-01, +9.82003689e-01},
	{-1.80002823e-01, +9.83843207e-01},
	{-1.70002818e-01, +9.85584319e-01},
	{-1.60002813e-01, +9.87226844e-01},
	{-1.50002807e-01, +9.88770664e-01},
	{-1.40002802e-01, +9.90215600e-01},
	{-1.30002797e-01, +9.91561532e-01},
	{-1.20002799e-01, +9.92808282e-01},
	{-1.10002801e-01, +9.93955791e-01},
	{-1.00002803e-01, +9.95003879e-01},
	{-9.00028050e-02, +9.95952487e-01},
	{-8.00028071e-02, +9.96801496e-01},
	{-7.00028092e-02, +9.97550786e-01},
	{-6.00028113e-02, +9.98200357e-01},
	{-5.00028133e-02, +9.98750091e-01},
	{-4.00028154e-02, +9.99199986e-01},
	{-3.00028156e-02, +9.99549925e-01},
	{-2.00028159e-02, +9.99799967e-01},
	{-1.00028161e-02, +9.99949992e-01},
	{-2.81631947e-06, +1.00000000e+00},
	{+9.99718346e-03, +9.99950051e-01},
	{+1.99971832e-02, +9.99800086e-01},
	{+2.99971830e-02, +9.99550104e-01},
	{+3.99971828e-02, +9.99200225e-01},
	{+4.99971807e-02, +9.98750389e-01},
	{+5.99971786e-02, +9.98200715e-01},
	{+6.99971765e-02, +9.97551203e-01},
	{+7.99971744e-02, +9.96801913e-01},
	{+8.99971724e-02, +9.95952964e-01},
	{+9.99971703e-02, +9.95004475e-01},
	{+1.09997168e-01, +9.93956387e-01},
	{+1.19997166e-01, +9.92808998e-01},
	{+1.29997164e-01, +9.91562247e-01},
	{+1.39997169e-01, +9.90216374e-01},
	{+1.49997175e-01, +9.88771498e-01},
	{+1.59997180e-01, +9.87227738e-01},
	{+1.69997185e-01, +9.85585272e-01},
	{+1.79997191e-01, +9.83844221e-01},
	{+1.89997196e-01, +9.82004762e-01},
	{+1.99997202e-01, +9.80067134e-01},
	{+2.09997207e-01, +9.78031516e-01},
	{+2.19997212e-01, +9.75898087e-01},
	{+2.29997218e-01, +9.73667026e-01},
	{+2.39997223e-01, +9.71338630e-01},
	{+2.49997228e-01, +9.68913078e-01},
	{+2.59997219e-01, +9.66390669e-01},
	{+2.69997209e-01, +9.63771641e-01},
	{+2.79997200e-01, +9.61056232e-01},
	{+2.89997190e-01, +9.58244681e-01},
	{+2.99997181e-01, +9.55337346e-01},
	{+3.09997171e-01, +9.52334404e-01},
	{+3.19997162e-01, +9.49236333e-01},
	{+3.29997152e-01, +9.46043253e-01},
	{+3.39997143e-01, +9.42755640e-01},
	{+3.49997133e-01, +9.39373672e-01},
	{+3.59997123e-01, +9.35897827e-01},
	{+3.69997114e-01, +9.32328403e-01},
	{+3.79997104e-01, +9.28665698e-01},
	{+3.89997095e-01, +9.24910188e-01},
	{+3.99997085e-01, +9.21062112e-01},
	{+4.09997076e-01, +9.17122006e-01},
	{+4.19997066e-01, +9.13090110e-01},
	{+4.29997057e-01, +9.08966959e-01},
	{+4.39997047e-01, +9.04752910e-01},
	{+4.49997038e-01, +9.00448382e-01},
	{+4.59997028e-01, +8.96053791e-01},
	{+4.69997019e-01, +8.91569614e-01},
	{+4.79997009e-01, +8.86996329e-01},
	{+4.89997000e-01, +8.82334292e-01},
	{+4.99996990e-01, +8.77583981e-01},
	{+5.09997010e-01, +8.72745991e-01},
	{+5.19997001e-01, +8.67820680e-01},
	{+5.29996991e-01, +8.62808585e-01},
	{+5.39996982e-01, +8.57710242e-01},
	{+5.49996972e-01, +8.52526128e-01},
	{+5.59996963e-01, +8.47256720e-01},
	{+5.69996953e-01, +8.41902614e-01},
	{+5.79996943e-01, +8.36464345e-01},
	{+5.89996934e-01, +8.30942392e-01},
	{+5.99996924e-01, +8.25337350e-01},
	{+6.09996915e-01, +8.19649756e-01},
	{+6.19996905e-01, +8.13880265e-01},
	{+6.29996896e-01, +8.08029354e-01},
	{+6.39996886e-01, +8.02097619e-01},
	{+6.49996877e-01, +7.96085715e-01},
	{+6.59996867e-01, +7.89994180e-01},
	{+6.69996858e-01, +7.83823609e-01},
	{+6.79996848e-01, +7.77574718e-01},
	{+6.89996839e-01, +7.71248043e-01},
	{+6.99996829e-01, +7.64844239e-01},
	{+7.09996819e-01, +7.58363962e-01},
	{+7.19996810e-01, +7.51807809e-01},
	{+7.29996800e-01, +7.45176554e-01},
	{+7.39996791e-01, +7.38470733e-01},
	{+7.49996781e-01, +7.31691062e-01},
	{+7.59996772e-01, +7.24838257e-01},
	{+7.69996762e-01, +7.17912912e-01},
	{+7.79996753e-01, +7.10915804e-01},
	{+7.89996743e-01, +7.03847647e-01},
	{+7.99996734e-01, +6.96709037e-01},
	{+8.09996724e-01, +6.89500809e-01},
	{+8.19996715e-01, +6.82223618e-01},
	{+8.29996705e-01, +6.74878180e-01},
	{+8.39996696e-01, +6.67465270e-01},
	{+8.49996686e-01, +6.59985662e-01},
	{+8.59996676e-01, +6.52440012e-01},
	{+8.69996667e-01, +6.44829094e-01},
	{+8.79996657e-01, +6.37153745e-01},
	{+8.89996648e-01, +6.29414618e-01},
	{+8.99996638e-01, +6.21612608e-01},
	{+9.09996629e-01, +6.13748431e-01},
	{+9.19996619e-01, +6.05822861e-01},
	{+9.29996610e-01, +5.97836673e-01},
	{+9.39996600e-01, +5.89790761e-01},
	{+9.49996591e-01, +5.81685841e-01},
	{+9.59996581e-01, +5.73522806e-01},
	{+9.69996572e-01, +5.65302372e-01},
	{+9.79996562e-01, +5.57025373e-01},
	{+9.89996552e-01, +5.48692763e-01},
	{+9.99996543e-01, +5.40305197e-01},
	{+1.00999653e+00, +5.31863630e-01},
	{+1.01999652e+00, +5.23368895e-01},
	{+1.02999651e+00, +5.14821827e-01},
	{+1.03999650e+00, +5.06223261e-01},
	{+1.04999650e+00, +4.97574091e-01},
	{+1.05999649e+00, +4.88875151e-01},
	{+1.06999648e+00, +4.80127335e-01},
	{+1.07999647e+00, +4.71331477e-01},
	{+1.08999646e+00, +4.62488502e-01},
	{+1.09999645e+00, +4.53599274e-01},
	{+1.10999644e+00, +4.44664717e-01},
	{+1.11999643e+00, +4.35685664e-01},
	{+1.12999642e+00, +4.26663041e-01},
	{+1.13999641e+00, +4.17597771e-01},
	{+1.14999640e+00, +4.08490717e-01},
	{+1.15999639e+00, +3.99342835e-01},
	{+1.16999638e+00, +3.90155017e-01},
	{+1.17999637e+00, +3.80928189e-01},
	{+1.18999636e+00, +3.71663243e-01},
	{+1.19999635e+00, +3.62361163e-01},
	{+1.20999634e+00, +3.53022814e-01},
	{+1.21999633e+00, +3.43649179e-01},
	{+1.22999632e+00, +3.34241182e-01},
	{+1.23999631e+00, +3.24799776e-01},
	{+1.24999630e+00, +3.15325856e-01},
	{+1.25999629e+00, +3.05820435e-01},
	{+1.26999629e+00, +2.96284407e-01},
	{+1.27999628e+00, +2.86718786e-01},
	{+1.28999627e+00, +2.77124465e-01},
	{+1.29999626e+00, +2.67502427e-01},
	{+1.30999625e+00, +2.57853657e-01},
	{+1.31999624e+00, +2.48179093e-01},
	{+1.32999623e+00, +2.38479719e-01},
	{+1.33999622e+00, +2.28756487e-01},
	{+1.34999621e+00, +2.19010383e-01},
	{+1.35999620e+00, +2.09242389e-01},
	{+1.36999619e+00, +1.99453458e-01},
	{+1.37999618e+00, +1.89644575e-01},
	{+1.38999617e+00, +1.79816738e-01},
	{+1.39999616e+00, +1.69970930e-01},
	{+1.40999615e+00, +1.60108104e-01},
	{+1.41999614e+00, +1.50229290e-01},
	{+1.42999613e+00, +1.40335441e-01},
	{+1.43999612e+00, +1.30427554e-01},
	{+1.44999611e+00, +1.20506629e-01},
	{+1.45999610e+00, +1.10573649e-01},
	{+1.46999609e+00, +1.00629620e-01},
	{+1.47999609e+00, +9.06755254e-02},
	{+1.48999608e+00, +8.07123631e-02},
	{+1.49999607e+00, +7.07411245e-02},
	{+1.50999606e+00, +6.07628189e-02},
	{+1.51999605e+00, +5.07784337e-02},
	{+1.52999604e+00, +4.07889709e-02},
	{+1.53999603e+00, +3.07954289e-02},
	{+1.54999602e+00, +2.07988080e-02},
	{+1.55999601e+00, +1.08001083e-02},
	{+1.56999600e+00, +8.00327398e-04},
	{+1.57999599e+00, -9.19953361e-03},
	{+1.58999598e+00, -1.91984735e-02},
	{+1.59999597e+00, -2.91954949e-02},
	{+1.60999596e+00, -3.91895957e-02},
	{+1.61999595e+00, -4.91797775e-02},
	{+1.62999594e+00, -5.91650419e-02},
	{+1.63999593e+00, -6.91443905e-02},
	{+1.64999592e+00, -7.91168213e-02},
	{+1.65999591e+00, -8.90813470e-02},
	{+1.66999590e+00, -9.90369618e-02},
	{+1.67999589e+00, -1.08982675e-01},
	{+1.68999588e+00, -1.18917480e-01},
	{+1.69999588e+00, -1.28840402e-01},
	{+1.70999587e+00, -1.38750434e-01},
	{+1.71999586e+00, -1.48646608e-01},
	{+1.72999585e+00, -1.58527896e-01},
	{+1.73999584e+00, -1.68393344e-01},
	{+1.74999583e+00, -1.78241953e-01},
	{+1.75999582e+00, -1.88072726e-01},
	{+1.76999581e+00, -1.97884709e-01},
	{+1.77999580e+00, -2.07676888e-01},
	{+1.78999579e+00, -2.17448309e-01},
	{+1.79999578e+00, -2.27197990e-01},
	{+1.80999577e+00, -2.36924946e-01},
	{+1.81999576e+00, -2.46628195e-01},
	{+1.82999575e+00, -2.56306797e-01},
	{+1.83999574e+00, -2.65959769e-01},
	{+1.84999573e+00, -2.75586158e-01},
	{+1.85999572e+00, -2.85184950e-01},
	{+1.86999571e+00, -2.94755250e-01},
	{+1.87999570e+00, -3.04296076e-01},
	{+1.88999569e+00, -3.13806474e-01},
	{+1.89999568e+00, -3.23285490e-01},
	{+1.90999568e+00, -3.32732171e-01},
	{+1.91999567e+00, -3.42145592e-01},
	{+1.92999566e+00, -3.51524770e-01},
	{+1.93999565e+00, -3.60868812e-01},
	{+1.94999564e+00, -3.70176792e-01},
	{+1.95999563e+00, -3.79447728e-01},
	{+1.96999562e+00, -3.88680726e-01},
	{+1.97999561e+00, -3.97874832e-01},
	{+1.98999560e+00, -4.07029182e-01},
	{+1.99999559e+00, -4.16142821e-01},
	{+2.00999570e+00, -4.25214946e-01},
	{+2.01999569e+00, -4.34244454e-01},
	{+2.02999568e+00, -4.43230540e-01},
	{+2.03999567e+00, -4.52172309e-01},
	{+2.04999566e+00, -4.61068839e-01},
	{+2.05999565e+00, -4.69919264e-01},
	{+2.06999564e+00, -4.78722721e-01},
	{+2.07999563e+00, -4.87478286e-01},
	{+2.08999562e+00, -4.96185124e-01},
	{+2.09999561e+00, -5.04842341e-01},
	{+2.10999560e+00, -5.13449013e-01},
	{+2.11999559e+00, -5.22004426e-01},
	{+2.12999558e+00, -5.30507624e-01},
	{+2.13999557e+00, -5.38957715e-01},
	{+2.14999557e+00, -5.47353983e-01},
	{+2.15999556e+00, -5.55695474e-01},
	{+2.16999555e+00, -5.63981354e-01},
	{+2.17999554e+00, -5.72210908e-01},
	{+2.18999553e+00, -5.80383241e-01},
	{+2.19999552e+00, -5.88497519e-01},
	{+2.20999551e+00, -5.96552908e-01},
	{+2.21999550e+00, -6.04548693e-01},
	{+2.22999549e+00, -6.12483978e-01},
	{+2.23999548e+00, -6.20358050e-01},
	{+2.24999547e+00, -6.28170073e-01},
	{+2.25999546e+00, -6.35919333e-01},
	{+2.26999545e+00, -6.43604934e-01},
	{+2.27999544e+00, -6.51226223e-01},
	{+2.28999543e+00, -6.58782363e-01},
	{+2.29999542e+00, -6.66272581e-01},
	{+2.30999541e+00, -6.73696220e-01},
	{+2.31999540e+00, -6.81052506e-01},
	{+2.32999539e+00, -6.88340664e-01},
	{+2.33999538e+00, -6.95560038e-01},
	{+2.34999537e+00, -7.02709794e-01},
	{+2.35999537e+00, -7.09789276e-01},
	{+2.36999536e+00, -7.16797829e-01},
	{+2.37999535e+00, -7.23734677e-01},
	{+2.38999534e+00, -7.30599165e-01},
	{+2.39999533e+00, -7.37390578e-01},
	{+2.40999532e+00, -7.44108260e-01},
	{+2.41999531e+00, -7.50751495e-01},
	{+2.42999530e+00, -7.57319689e-01},
	{+2.43999529e+00, -7.63812184e-01},
	{+2.44999528e+00, -7.70228267e-01},
	{+2.45999527e+00, -7.76567280e-01},
	{+2.46999526e+00, -7.82828689e-01},
	{+2.47999525e+00, -7.89011836e-01},
	{+2.48999524e+00, -7.95116067e-01},
	{+2.49999523e+00, -8.01140785e-01},
	{+2.50999522e+00, -8.07085335e-01},
	{+2.51999521e+00, -8.12949240e-01},
	{+2.52999520e+00, -8.18731844e-01},
	{+2.53999519e+00, -8.24432552e-01},
	{+2.54999518e+00, -8.30050826e-01},
	{+2.55999517e+00, -8.35586131e-01},
	{+2.56999516e+00, -8.41037869e-01},
	{+2.57999516e+00, -8.46405447e-01},
	{+2.58999515e+00, -8.51688445e-01},
	{+2.59999514e+00, -8.56886268e-01},
	{+2.60999513e+00, -8.61998379e-01},
	{+2.61999512e+00, -8.67024302e-01},
	{+2.62999511e+00, -8.71963501e-01},
	{+2.63999510e+00, -8.76815557e-01},
	{+2.64999509e+00, -8.81579876e-01},
	{+2.65999508e+00, -8.86256039e-01},
	{+2.66999507e+00, -8.90843630e-01},
	{+2.67999506e+00, -8.95342112e-01},
	{+2.68999505e+00, -8.99751067e-01},
	{+2.69999504e+00, -9.04070020e-01},
	{+2.70999503e+00, -9.08298612e-01},
	{+2.71999502e+00, -9.12436306e-01},
	{+2.72999501e+00, -9.16482806e-01},
	{+2.73999500e+00, -9.20437694e-01},
	{+2.74999499e+00, -9.24300492e-01},
	{+2.75999498e+00, -9.28070843e-01},
	{+2.76999497e+00, -9.31748390e-01},
	{+2.77999496e+00, -9.35332775e-01},
	{+2.78999496e+00, -9.38823640e-01},
	{+2.79999495e+00, -9.42220628e-01},
	{+2.80999494e+00, -9.45523381e-01},
	{+2.81999493e+00, -9.48731601e-01},
	{+2.82999492e+00, -9.51844931e-01},
	{+2.83999491e+00, -9.54863131e-01},
	{+2.84999490e+00, -9.57785785e-01},
	{+2.85999489e+00, -9.60612655e-01},
	{+2.86999488e+00, -9.63343501e-01},
	{+2.87999487e+00, -9.65977967e-01},
	{+2.88999486e+00, -9.68515873e-01},
	{+2.89999485e+00, -9.70956922e-01},
	{+2.90999484e+00, -9.73300874e-01},
	{+2.91999483e+00, -9.75547493e-01},
	{+2.92999482e+00, -9.77696598e-01},
	{+2.93999481e+00, -9.79747891e-01},
	{+2.94999480e+00, -9.81701195e-01},
	{+2.95999479e+00, -9.83556390e-01},
	{+2.96999478e+00, -9.85313177e-01},
	{+2.97999477e+00, -9.86971438e-01},
	{+2.98999476e+00, -9.88531053e-01},
	{+2.99999475e+00, -9.89991784e-01},
	{+3.00999475e+00, -9.91353512e-01},
	{+3.01999474e+00, -9.92616057e-01},
	{+3.02999473e+00, -9.93779421e-01},
	{+3.03999472e+00, -9.94843364e-01},
	{+3.04999471e+00, -9.95807827e-01},
	{+3.05999470e+00, -9.96672750e-01},
	{+3.06999469e+00, -9.97437954e-01},
	{+3.07999468e+00, -9.98103440e-01},
	{+3.08999467e+00, -9.98669147e-01},
	{+3.09999466e+00, -9.99134958e-01},
	{+3.10999465e+00, -9.99500811e-01},
	{+3.11999464e+00, -9.99766767e-01},
	{+3.12999463e+00, -9.99932766e-01},
	{+3.13999462e+00, -9.99998748e-01},
}

//------------------------------------------------------------------------------
