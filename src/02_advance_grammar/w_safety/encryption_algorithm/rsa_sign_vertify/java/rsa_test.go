// @Time : 2023/3/20 18:00
// @Author : xiaoweiwei
// @File : rsa_test

package java

import "testing"

// RSACryptTest RSA加密参数信息
type RSACryptTest struct {
	data       string
	encodeType Encode
	hashType   Hash
}

// 格式：PKCS8
// 密钥长度位数：2048
func TestRSA(t *testing.T) {
	handle := NewRSACrypt(RSASecret{
		PublicKey:          "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvtyZ8TWGbZRIzo7Rx1Pk\nW+fN9mVVK0xyxZKL11wwrmXROCEL78fWkXmctD76YR4eQzHZw9gw2yTpbXRUbDt9\nNzhtEnsFaa+HTcbpQpumd5GZ1EXxIxAp1wuVQ4+LXT2aypBoBuxFwmXEpF9nMrqn\ngj3acwq9LOR9UIpoxP1xyUtpidPPFfqP0Pbbc+cyBQl79e5vRfgPXoUvfwMTMZlm\nGJP8QNk+x46hgjJlR+q7l21QJok9cMzzHADXNSOb3ep9nQqmFD3zui0IKK2BHcND\n5srn7qXo5Qz85I0V8FGKvAT0DDtBXhfhcJk4pq2W9UXFjMmAqk31iREZqEadua0K\nXwIDAQAB",
		PublicKeyDataType:  Base64,
		PrivateKey:         "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC+3JnxNYZtlEjO\njtHHU+Rb5832ZVUrTHLFkovXXDCuZdE4IQvvx9aReZy0PvphHh5DMdnD2DDbJOlt\ndFRsO303OG0SewVpr4dNxulCm6Z3kZnURfEjECnXC5VDj4tdPZrKkGgG7EXCZcSk\nX2cyuqeCPdpzCr0s5H1QimjE/XHJS2mJ088V+o/Q9ttz5zIFCXv17m9F+A9ehS9/\nAxMxmWYYk/xA2T7HjqGCMmVH6ruXbVAmiT1wzPMcANc1I5vd6n2dCqYUPfO6LQgo\nrYEdw0PmyufupejlDPzkjRXwUYq8BPQMO0FeF+FwmTimrZb1RcWMyYCqTfWJERmo\nRp25rQpfAgMBAAECggEBALqw/B+2FIURO9pYVxhblg7UAmIWoqWUxNs/Smxg7/lW\n/LyjjpJ/7gEkbMTeqXjfxghCNdlcmgVx6Ka42G96DTQ7jG08SoIjQP0yKS6PI9jn\nYgoCVCnRAKwp11pbiZ2tycCEWk0TRfUuKJgD/eTgMUOzjv9irMLByVZD3NF7e5bB\n62uYcAtpZg+lzsZc3K8mZG01tIb/Wsc/RKH/cO84hPWT+NM1uBlYP8cXAJ5fcSBg\nIx49z2NeOe+s0mIYBy1Z8xAXz622Zuh+v314V6jnAFIvGnpp//Os1V2e6K64DfsX\nVv979kZh01BmlkZNsCioWR7+kpPVqOLuItq0hUUSHikCgYEA9X8W4hqGe0jAaPD3\nzyoY+z9my/LAq8wW5i5c9ZsPDUf40w5Skvba6Wkj0hCAyeT868vwMVUnfeqTaUXH\nslx3kJh/a0vK31+4bisg2Bgnwaya1SInTFOEVr0ABcVa5esNHcjI4uTLK+Wj1xz0\nMpwgb/X2K1bsuBrEpCdf62mVjPUCgYEAxwcZvCGIfy9Wyd8X5kWGOAaeL6DJG1YK\nrJ56FoyH8TFGmRneS04prVeYX8iJF03Wyj/MUAs4BCCBEq5ZGYNs2J55z6UKD7KI\nSAMVxVGJRxh3JWwUL0sGvO5qaSnOcGCy8xrXbqRVKBDmL+6K/YHv181acPyX4SLL\nYS98OCWcpYMCgYBpaA8AN9lkF8GdjVEpXu7o+bj+epVlbjSq0l9RiSk2T5+zpyOn\nPFyT+XPK3xEdHrfF2oOf+usA8nYmrjJnD7K6n++dtyY9MJE0pEnu3rg7PQwh1Q9A\ng1+ACph24dz8eSbhkfNHEGt3xprJYwIqtMRsrhPrWeA70Cp8PPe/UU9H/QKBgDlD\nabwfSc6OKEJc1duDFWL5RHofoL0kvF1+G1JxjItwygcy0iqIiqU9FrI4WHlEBBCc\n2oQNNki5sWlWkHwuvBp3PkToD0UE8QEHIjRriTvTZjR2LDiPNZkWX7dyN9tNdy9L\nAcDrtzSX6CLSV1spD4DfrK3lfy1ffXSw3OSqra9RAoGBAPIApsZG3KYzISbwYg8W\njqbZzMDkYnYaMrClzY7f59nbOFa90jctt00okeav1xKUZyHLZKbP/pcbVhGPZCaT\nA8CbILd6zGje23JknWdYV1APax5l9h0jryr5BWJfKRUypXKD4XLHHZxZMNqwZrZr\npyPL8R2T+CRweEtXnKQaRR2o",
		PrivateKeyType:     PKCS8,
		PrivateKeyDataType: Base64,
	})
	var tests = []RSACryptTest{
		{
			"测试加密数据",
			Base64,
			MD5,
		},
	}

	for _, v := range tests {
		println("====== start ===== ")
		println("rsaTest encodeType ", v.encodeType, "  hashType is ", v.hashType)
		//encrypt data & encode result
		println(" data is  " + v.data)
		encrypt, err := handle.Encrypt(v.data, v.encodeType)
		if err != nil {
			println("encrypt error : %v", err)
			return
		}
		println("encrypt data is ", encrypt)

		// decrypt encrypted & encoded data
		decrypt, err := handle.Decrypt(encrypt, v.encodeType)
		if err != nil {
			println("decrypt error : %v", err)
			return
		}
		println("decrypt data is ", decrypt)

		//sign data with digest algorithm & encode result
		sign, err := handle.Sign(v.data, v.hashType, v.encodeType)
		if err != nil {
			t.Fatalf("sign error : %v", err)
		}
		println("sign data is ", sign)

		//verify data that signed with digest algorithm & encoded whether match original data
		verifySign, err := handle.VerifySign(v.data, v.hashType, sign, v.encodeType)
		if err != nil {
			t.Fatalf("verifySign error : %v", err)
		}
		if !verifySign {
			t.Fatal("verifySign result failed")
		}
		println("verifySign result is ", verifySign)

		println("====== end ===== ")

	}

}
