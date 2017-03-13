package mtproto

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

const randomness = `
3E0549828CCA27E966B301A48FECE2FC
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D

311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D
311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC02

6F620AFA575C9233EB4C014110A7BCAF49464F798A18A0981FEA1E05E8DA67D9681E0FD6DF0EDF0272AE3492451A84502F2EFC0DA18741A5FB80BD82296919A70FAA6D07CBBBCA2037EA7D3E327B61D585ED3373EE0553A91CBD29B01FA9A89D479CA53D57BDE3A76FBD922A923A0A38B922C1D0701F53FF52D7EA9217080163A64901E766EB6A0F20BC391B64B9D1DD2CD13A7D0C946A3A7DF8CEC9E2236446F646C42CFE2B60A2A8D776E56C8D7519B08B88ED0970E10D12A8C9E355D765F2B7BBB7B4CA9360083435523CB0D57D2B106FD14F94B4EEE79D8AC131CA56AD389C84FE279716F8124A543337FB9EA3D988EC5FA63D90A4BA3970E7A39E5C0DE5

311C85DB234AA2640AFC4A76
`

// example requests and responses from https://core.telegram.org/mtproto/samples-auth_key

const req1 = `
00 00 00 00 00 00 00 00 4A 96 70 27 C4 7A E5 51
14 00 00 00 78 97 46 60 3E 05 49 82 8C CA 27 E9
66 B3 01 A4 8F EC E2 FC
`

const res1 = `
00 00 00 00 00 00 00 00 01 C8 83 1E C9 7A E5 51
40 00 00 00 63 24 16 05 3E 05 49 82 8C CA 27 E9
66 B3 01 A4 8F EC E2 FC A5 CF 4D 33 F4 A1 1E A8
77 BA 4A A5 73 90 73 30 08 17 ED 48 94 1A 08 F9
81 00 00 00 15 C4 B5 1C 01 00 00 00 21 6B E8 6C
02 2B B4 C3`

const req2 = `
00 00 00 00 00 00 00 00 27 7A 71 17 C9 7A E5 51
40 01 00 00 BE E4 12 D7 3E 05 49 82 8C CA 27 E9
66 B3 01 A4 8F EC E2 FC A5 CF 4D 33 F4 A1 1E A8
77 BA 4A A5 73 90 73 30 04 49 4C 55 3B 00 00 00
04 53 91 10 73 00 00 00 21 6B E8 6C 02 2B B4 C3
FE 00 01 00 7B B0 10 0A 52 31 61 90 4D 9C 69 FA
04 BC 60 DE CF C5 DD 74 B9 99 95 C7 68 EB 60 D8
71 6E 21 09 BA F2 D4 60 1D AB 6B 09 61 0D C1 10
67 BB 89 02 1E 09 47 1F CF A5 2D BD 0F 23 20 4A
D8 CA 8B 01 2B F4 0A 11 2F 44 69 5A B6 C2 66 95
53 86 11 4E F5 21 1E 63 72 22 7A DB D3 49 95 D3
E0 E5 FF 02 EC 63 A4 3F 99 26 87 89 62 F7 C5 70
E6 A6 E7 8B F8 36 6A F9 17 A5 27 26 75 C4 60 64
BE 62 E3 E2 02 EF A8 B1 AD FB 1C 32 A8 98 C2 98
7B E2 7B 5F 31 D5 7C 9B B9 63 AB CB 73 4B 16 F6
52 CE DB 42 93 CB B7 C8 78 A3 A3 FF AC 9D BE A9
DF 7C 67 BC 9E 95 08 E1 11 C7 8F C4 6E 05 7F 5C
65 AD E3 81 D9 1F EE 43 0A 6B 57 6A 99 BD F8 55
1F DB 1B E2 B5 70 69 B1 A4 57 30 61 8F 27 42 7E
8A 04 72 0B 49 71 EF 4A 92 15 98 3D 68 F2 83 0C
3E AA 6E 40 38 55 62 F9 70 D3 8A 05 C9 F1 24 6D
C3 34 38 E6
`

const res2 = `
00 00 00 00 00 00 00 00 01 54 43 36 CB 7A E5 51
78 02 00 00 5C 07 E8 D0 3E 05 49 82 8C CA 27 E9
66 B3 01 A4 8F EC E2 FC A5 CF 4D 33 F4 A1 1E A8
77 BA 4A A5 73 90 73 30 FE 50 02 00 28 A9 2F E2
01 73 B3 47 A8 BB 32 4B 5F AB 26 67 C9 A8 BB CE
64 68 D5 B5 09 A4 CB DD C1 86 24 0A C9 12 CF 70
06 AF 89 26 DE 60 6A 2E 74 C0 49 3C AA 57 74 1E
6C 82 45 1F 54 D3 E0 68 F5 CC C4 9B 44 44 12 4B
96 66 FF B4 05 AA B5 64 A3 D0 1E 67 F6 E9 12 86
7C 8D 20 D9 88 27 07 DC 33 0B 17 B4 E0 DD 57 CB
53 BF AA FA 9E F5 BE 76 AE 6C 1B 9B 6C 51 E2 D6
50 2A 47 C8 83 09 5C 46 C8 1E 3B E2 5F 62 42 7B
58 54 88 BB 3B F2 39 21 3B F4 8E B8 FE 34 C9 A0
26 CC 84 13 93 40 43 97 4D B0 35 56 63 30 38 39
2C EC B5 1F 94 82 4E 14 0B 98 63 77 30 A4 BE 79
A8 F9 DA FA 39 BA E8 1E 10 95 84 9E A4 C8 34 67
C9 2A 3A 17 D9 97 81 7C 8A 7A C6 1C 3F F4 14 DA
37 B7 D6 6E 94 9C 0A EC 85 8F 04 82 24 21 0F CC
61 F1 1C 3A 91 0B 43 1C CB D1 04 CC CC 8D C6 D2
9D 4A 5D 13 3B E6 39 A4 C3 2B BF F1 53 E6 3A CA
3A C5 2F 2E 47 09 B8 AE 01 84 4B 14 2C 1E E8 9D
07 5D 64 F6 9A 39 9F EB 04 E6 56 FE 36 75 A6 F8
F4 12 07 8F 3D 0B 58 DA 15 31 1C 1A 9F 8E 53 B3
CD 6B B5 57 2C 29 49 04 B7 26 D0 BE 33 7E 2E 21
97 7D A2 6D D6 E3 32 70 25 1C 2C A2 9D FC C7 02
27 F0 75 5F 84 CF DA 9A C4 B8 DD 5F 84 F1 D1 EB
36 BA 45 CD DC 70 44 4D 8C 21 3E 4B D8 F6 3B 8A
B9 5A 2D 0B 41 80 DC 91 28 3D C0 63 AC FB 92 D6
A4 E4 07 CD E7 C8 C6 96 89 F7 7A 00 74 41 D4 A6
A8 38 4B 66 65 02 D9 B7 7F C6 8B 5B 43 CC 60 7E
60 A1 46 22 3E 11 0F CB 43 BC 3C 94 2E F9 81 93
0C DC 4A 1D 31 0C 0B 64 D5 E5 5D 30 8D 86 32 51
AB 90 50 2C 3E 46 CC 59 9E 88 6A 92 7C DA 96 3B
9E B1 6C E6 26 03 B6 85 29 EE 98 F9 F5 20 64 19
E0 3F B4 58 EC 4B D9 45 4A A8 F6 BA 77 75 73 CC
54 B3 28 89 5B 1D F2 5E AD 9F B4 CD 51 98 EE 02
2B 2B 81 F3 88 D2 81 D5 E5 BC 58 01 07 CA 01 A5
06 65 C3 2B 55 27 15 F3 35 FD 76 26 4F AD 00 DD
D5 AE 45 B9 48 32 AC 79 CE 7C 51 1D 19 4B C4 2B
70 EF A8 50 BB 15 C2 01 2C 52 15 CA BF E9 7C E6
6B 8D 87 34 D0 EE 75 9A 63 8A F0 13
`

const req3 = `
00 00 00 00 00 00 00 00 6D 2C A3 2A CD 7A E5 51
78 01 00 00 1F 5F 04 F5 3E 05 49 82 8C CA 27 E9
66 B3 01 A4 8F EC E2 FC A5 CF 4D 33 F4 A1 1E A8
77 BA 4A A5 73 90 73 30 FE 50 01 00 92 8A 49 57
D0 46 3B 52 5C 1C C4 8A AB AA 03 0A 25 6B E5 C7
46 79 2C 84 CA 4C 5A 0D F6 0A C7 99 04 8D 98 A3
8A 84 80 ED CF 08 22 14 DF C7 9D CB 9E E3 4E 20
65 13 E2 B3 BC 15 04 CF E6 C9 AD A4 6B F9 A0 3C
A7 4F 19 2E AF 8C 27 84 54 AD AB C7 95 A5 66 61
54 62 D3 18 17 38 29 84 03 95 05 F7 1C B3 3A 41
E2 52 7A 4B 1A C0 51 07 87 2F ED 8E 3A BC EE 15
18 AE 96 5B 0E D3 AE D7 F6 74 79 15 5B DA 8E 4C
28 6B 64 CD F1 23 EC 74 8C F2 89 B1 DB 02 D1 90
7B 56 2D F4 62 D8 58 2B A6 F0 A3 02 2D C2 D3 50
4D 69 D1 BA 48 B6 77 E3 A8 30 BF AF D6 75 84 C8
AA 24 E1 34 4A 89 04 E3 05 F9 58 7C 92 EF 96 4F
00 83 F5 0F 61 EA B4 A3 93 EA A3 3C 92 70 29 4A
ED C7 73 28 91 D4 EA 15 99 F5 23 11 D7 44 69 D2
11 2F 4E DF 3F 34 2E 93 C8 E8 7E 81 2D C3 98 9B
AE CF E6 74 0A 46 07 75 24 C7 50 93 F5 A5 40 57
36 DE 89 37 BB 6E 42 C9 A0 DC F2 2C A5 32 27 D4
62 BC CC 2C FE 94 B6 FE 86 AB 7F BF A3 95 02 1F
66 66 1A F7 C0 02 4C A2 98 6C A0 3F 34 76 90 54
07 D1 EA 9C 01 0B 76 32 58 DB 1A A2 CC 78 26 D9
13 34 EF C1 FD C6 65 B6 7F E4 5E D0
`

const res3 = `
00 00 00 00 00 00 00 00 01 30 AA C5 CE 7A E5 51
34 00 00 00 34 F7 CB 3B 3E 05 49 82 8C CA 27 E9
66 B3 01 A4 8F EC E2 FC A5 CF 4D 33 F4 A1 1E A8
77 BA 4A A5 73 90 73 30 CC EB C0 21 72 66 E1 ED
EC 7F B0 A0 EE D6 C2 20
`

const expectedKeyStr = `AB96E207C631300986F30EF97DF55E179E63C112675F0CE502EE76D74BBEE6CBD1E95772818881E9F2FF54BD52C258787474F6A7BEA61EABE49D1D01D55F64FC07BC31685716EC8FB46FEACF9502E42CFD6B9F45A08E90AA5C2B5933AC767CBE1CD50D8E64F89727CA4A1A5D32C0DB80A9FCDBDDD4F8D5A1E774198F1A4299F927C484FEEC395F29647E43C3243986F93609E23538C21871DF50E00070B3B6A8FA9BC15628E8B43FF977409A61CEEC5A21CF7DFB5A4CC28F5257BC30CD8F2FB92FBF21E28924065F50E0BBD5E11A420300E2C136B80E9826C6C5609B5371B7850AA628323B6422F3A94F6DFDE4C3DC1EA60F7E11EE63122B3F39CBD1A8430157`

func TestKeyExchange(t *testing.T) {
	var keyex KeyEx
	var framer Framer
	var err error

	keyex.RandomReader = bytes.NewReader(fromHex(randomness))
	keyex.PubKey, err = ParsePublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}

	// --- req 1

	framer.MsgIDOverride = 0x51e57ac42770964a
	msgbytes, err := framer.Format(keyex.Start())
	if err != nil {
		t.Fatal(err)
	}

	a, e := hex.EncodeToString(msgbytes), hex.EncodeToString(fromHex(req1))
	if a != e {
		t.Errorf("req_pq is %q, expected %q", a, e)
	}

	// --- res 1

	inmsg, err := framer.Parse(fromHex(res1))
	if err != nil {
		t.Fatal(err)
	}
	msg, err := keyex.Handle(NewReader(inmsg.Payload))
	if err != nil {
		t.Fatal(err)
	}

	// --- req 2

	if msg == nil {
		t.Fatal("no reply to res_pq")
	}
	framer.MsgIDOverride = 0x51e57ac917717a27
	msgbytes, err = framer.Format(*msg)
	if err != nil {
		t.Fatal(err)
	}
	emsgbytes := fromHex(req2)
	a, e = hex.EncodeToString(msgbytes), hex.EncodeToString(emsgbytes)
	if len(msgbytes) != len(emsgbytes) {
		t.Errorf("req_DH_params is %v, expected %v (len mismatch: got %v, wanted %v)", a, e, len(msgbytes), len(emsgbytes))
	}
	// if a != e {
	// 	t.Errorf("req_DH_params is %v, expected %v", a, e)
	// }

	// --- res 2

	inmsg, err = framer.Parse(fromHex(res2))
	if err != nil {
		t.Fatal(err)
	}
	msg, err = keyex.Handle(NewReader(inmsg.Payload))
	if err != nil {
		t.Fatal(err)
	}

	// --- req 3

	if msg == nil {
		t.Fatal("no reply to server_DH_params_ok")
	}
	framer.MsgIDOverride = 0x51e57acd2aa32c6d
	msgbytes, err = framer.Format(*msg)
	if err != nil {
		t.Fatal(err)
	}
	emsgbytes = fromHex(req3)
	a, e = hex.EncodeToString(msgbytes), hex.EncodeToString(emsgbytes)
	if len(msgbytes) != len(emsgbytes) {
		t.Errorf("set_client_DH_params is %v, expected %v (len mismatch: got %v, wanted %v)", a, e, len(msgbytes), len(emsgbytes))
	}

	// --- res 3

	inmsg, err = framer.Parse(fromHex(res3))
	if err != nil {
		t.Fatal(err)
	}
	msg, err = keyex.Handle(NewReader(inmsg.Payload))
	if err != nil {
		t.Fatal(err)
	}

	// --- done

	auth, err := keyex.Result()
	if err != nil {
		t.Fatal(err)
	}

	expectedKey := fromHex(expectedKeyStr)
	if !bytes.Equal(auth.Key, expectedKey) {
		t.Errorf("key is %x, expected %x", auth.Key, expectedKey)
	}
}

func fromHex(s string) []byte {
	data, err := hex.DecodeString(strings.Map(dropSpace, s))
	if err != nil {
		panic(err)
	}
	return data
}

func dropSpace(r rune) rune {
	if r == ' ' || r == '\t' || r == '\n' {
		return -1
	} else {
		return r
	}
}
