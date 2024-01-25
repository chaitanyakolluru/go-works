package testterraform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

// NowProcessAndTest function takes in the generated plan.out file and tests with it.
func NowProcessAndTest(cwd, testcondition string) {

	tfState, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + ".temp" + string(os.PathSeparator) + "terraform.tfstate")
	mapTfStateOld := make(map[string]interface{})
	if err := json.Unmarshal([]byte(tfState), &mapTfStateOld); err != nil {
		log.Fatal(err)
	}

	obj := make(map[interface{}]interface{})
	mapTfState := make(map[interface{}]interface{})
	for l, m := range mapTfStateOld {
		mapTfState[l] = m
	}

	for _, vI := range mapTfState["resources"].([]interface{}) {

		vII := vI
		if len(vII.(map[string]interface{})["instances"].([]interface{})) != 0 {
			for _, vIV := range vII.(map[string]interface{})["instances"].([]interface{}) {

				vVI := vIV.(map[string]interface{})["attributes"].(map[string]interface{})
				vVtoUse := make(map[interface{}]interface{})
				for z, zi := range vVI {
					switch reflect.ValueOf(zi).Kind() {
					case reflect.String, reflect.Bool:
						vVtoUse[z] = zi
					case reflect.Slice:
						vVtoUse[z] = zi.([]interface{})
					case reflect.Map:
						lk := make(map[interface{}]interface{})
						for kk, ll := range zi.(map[string]interface{}) {
							lk[kk] = ll
						}
						vVtoUse[z] = lk
					}
				}
				if _, ok := vVI["name"]; ok {

					obj[vVI["name"]] = vVtoUse
				}
			}
		}
	}

	testCon, _ := ioutil.ReadFile(cwd + string(os.PathSeparator) + testcondition)
	mapTestCon := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(testCon), &mapTestCon); err != nil {
		log.Fatal(err)
	}

	// nowCompareMain(obj, mapTestCon, resConRes)
	objC := prepareCondensedObj(obj)

	compareObjCmapTestCon(objC, mapTestCon)

	// writeObjCondensed(objC, cwd)

}

// writeObjCondensed function writes objCondensed as yaml file.
func writeObjCondensed(objC map[interface{}]interface{}, cwd string) {

	objCP := make(map[string]interface{})
	for obk, obv := range objC {
		obvI := make(map[string][]interface{})
		for obvi, obvii := range obv.(map[interface{}][]interface{}) {
			obvI[fmt.Sprint(obvi)] = obvii
		}
		objCP[fmt.Sprint(obk)] = obvI
	}
	objCPrint, err := yaml.Marshal(objCP)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(cwd+string(os.PathSeparator)+"objCondensed.yaml", []byte(objCPrint), 0644); err != nil {
		log.Fatal(err)
	}

}

// compareObjCmapTestCon function compares test inputted yaml data with objCondensed.
func compareObjCmapTestCon(o map[interface{}]interface{}, m map[interface{}]interface{}) {

	r := make(map[interface{}]interface{})
	for km, vm := range m {

		switch vm.(type) {
		case map[interface{}]interface{}:
			if _, ok := o[km]; ok {
				r[km] = make(map[interface{}]interface{})
				for kmi, vmi := range vm.(map[interface{}]interface{}) {

					for _, voi := range o[km].(map[interface{}][]interface{})[kmi] {

						if reflect.ValueOf(vmi).Kind() == reflect.Int {
							if reflect.ValueOf(voi).Kind() == reflect.Float64 && int(voi.(float64)) == vmi || voi == vmi {
								if _, okR := r[km]; okR {
									fmt.Printf("%v:%v:%v :: TRUE\n", km, kmi, vmi)
									r[km].(map[interface{}]interface{})[kmi] = vmi
								}
							}
						} else if reflect.ValueOf(vmi).Kind() == reflect.Float64 {
							if reflect.ValueOf(voi).Kind() == reflect.Int && float64(voi.(int)) == vmi || voi == vmi {
								if _, okR := r[km]; okR {
									fmt.Printf("%v:%v:%v :: TRUE\n", km, kmi, vmi)
									r[km].(map[interface{}]interface{})[kmi] = vmi
								}
							}
						} else {
							if voi == vmi {
								if _, okR := r[km]; okR {
									fmt.Printf("%v:%v:%v :: TRUE\n", km, kmi, vmi)
									r[km].(map[interface{}]interface{})[kmi] = vmi
								}

							}
						}

					}

				}
			}
		case nil:
			if _, ok := o[km]; ok {
				fmt.Printf("%v :: TRUE\n", km)
				r[km] = make(map[interface{}]interface{})
			}
		}

	}

	fmt.Println("=======================================================================================")
	passed := 1
	for km, vm := range m {
		switch vm.(type) {
		case nil:
			if _, ok := r[km]; !ok {
				passed = 0
				fmt.Printf("%v :: FALSE\n", km)
			}
		case map[interface{}]interface{}:
			for kmi, vmi := range vm.(map[interface{}]interface{}) {
				if _, ok := r[km]; ok {
					if vmi != r[km].(map[interface{}]interface{})[kmi] {
						passed = 0
						fmt.Printf("%v:%v:%v :: FALSE\n", km, kmi, vmi)
					}
				} else {
					passed = 0
					fmt.Printf("%v:%v:%v :: FALSE\n", km, kmi, vmi)
				}
			}
		}

	}
	if passed == 0 {
		fmt.Println("==============================TESTS FAILED=============================================")
	} else {
		fmt.Println("==============================TESTS PASSED=============================================")
	}
}

func processAllData(koi, voi interface{}, oCKo map[interface{}][]interface{}) map[interface{}][]interface{} {
	switch voi.(type) {
	case int, bool, string, float64:
		oCKo[koi] = append(oCKo[koi], voi)
	case map[interface{}]interface{}:
		for koii, voii := range voi.(map[interface{}]interface{}) {
			processAllData(koii, voii, oCKo)
		}
	case []interface{}:
		for _, voiii := range voi.([]interface{}) {
			processAllData(koi, voiii, oCKo)
		}
	case map[string]interface{}:
		for koii, voii := range voi.(map[string]interface{}) {
			processAllData(koii, voii, oCKo)
		}
	}
	return oCKo
}

func prepareCondensedObj(o map[interface{}]interface{}) map[interface{}]interface{} {

	oC := make(map[interface{}]interface{})
	for ko, vo := range o {
		oC[ko] = make(map[interface{}][]interface{})
		for koi, voi := range vo.(map[interface{}]interface{}) {
			oC[ko] = processAllData(koi, voi, oC[ko].(map[interface{}][]interface{}))
		}

	}

	return oC

}

// nowCompareMain function takes in two map interfaces from obj and testcondition and compares elements.
func nowCompareMain(o, m map[interface{}]interface{}, r map[string]interface{}) {

	for k, v := range m {
		if v == nil {

			if _, ok := o[k]; ok {
				fmt.Println(k, ":: TRUE")
			} else {
				fmt.Println("nil", k, "false")
				r[fmt.Sprint(k)] = v
			}

		} else if reflect.ValueOf(v).Kind() == reflect.String || reflect.ValueOf(v).Kind() == reflect.Bool {
			if o[k] == m[k] {
				fmt.Println(k, "=", m[k], ":: TRUE")
			} else {
				fmt.Println(v, reflect.ValueOf(v).Kind(), o[k], reflect.ValueOf(o[k]).Kind(), k, "false")
				r[fmt.Sprint(k)] = v
			}
		} else if reflect.ValueOf(v).Kind() == reflect.Float64 || reflect.ValueOf(v).Kind() == reflect.Int {
			if reflect.ValueOf(o[k]).Kind() == reflect.Float64 && int(o[k].(float64)) == m[k] || o[k] == m[k] {
				fmt.Println(k, "=", m[k], ":: TRUE")
			} else {
				fmt.Println(v, reflect.ValueOf(v).Kind(), o[k], reflect.ValueOf(o[k]).Kind(), k, "false")
				r[fmt.Sprint(k)] = v
			}
		} else if reflect.ValueOf(v).Kind() == reflect.Map {

			nowCompareMain(o[k].(map[interface{}]interface{}), m[k].(map[interface{}]interface{}), r)

		} else if reflect.ValueOf(v).Kind() == reflect.Slice {

			for xii, vii := range m[k].([]interface{}) {
				x := make(map[interface{}]interface{})
				for xii, yii := range o[k].([]interface{})[xii].(map[string]interface{}) {
					x[xii] = yii
				}
				nowCompareMain(x, vii.(map[interface{}]interface{}), r)
			}

		}

	}

}

// nowCompare function first form
func nowCompare(o, m map[interface{}]interface{}) {

	for k, v := range m {
		if v == nil {

			if _, ok := o[k]; ok {
				fmt.Println(k, "true")
			}

		} else if reflect.ValueOf(v).Kind() == reflect.String || reflect.ValueOf(v).Kind() == reflect.Bool {
			if o[k] == m[k] {
				fmt.Println("string", o[k], m[k], k, "true")
			}
		} else if reflect.ValueOf(v).Kind() == reflect.Map {

			x := make(map[interface{}]interface{})

			for xi, yi := range o[k].(map[interface{}]interface{}) {
				x[xi] = yi
			}
			y := make(map[interface{}]interface{})
			for xii, yii := range m[k].(map[interface{}]interface{}) {
				y[xii] = yii
			}

			nowCompare(x, y)

		} else if reflect.ValueOf(v).Kind() == reflect.Slice {

			for xii, vii := range m[k].([]interface{}) {
				nowCompare(v.([]interface{})[xii].(map[interface{}]interface{}), vii.(map[interface{}]interface{}))
			}

		}

	}

}
