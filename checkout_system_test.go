package code_kata

import "testing"

func TestAnalyzePricingRules(t *testing.T) {
	for _, c := range []struct {
		want_name  string
		want_price int
		s_price    int
		s_quantity int
		file_path  string
	}{
		{"A", 50, 130, 3, "tmp/pricing_rules.txt"},
	} {
		got := analyzePricingRules(c.file_path)

		if got[0].name != c.want_name || got[0].price != c.want_price || got[0].specialRule.price != c.s_price || got[0].specialRule.quantity != c.s_quantity {
			t.Errorf("test failed, expected %s, %d, %d, %d,  got %s, %d, %d, %d", c.want_name, c.want_price, c.s_price, c.s_quantity, got[0].name, got[0].price, got[0].specialRule.price, got[0].specialRule.quantity)
		}
	}
}
