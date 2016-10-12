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

		if got["A"].name != c.want_name || got["A"].price != c.want_price || got["A"].specialRule.price != c.s_price || got["A"].specialRule.quantity != c.s_quantity {
			t.Errorf("test failed, expected %s, %d, %d, %d,  got %s, %d, %d, %d", c.want_name, c.want_price, c.s_price, c.s_quantity, got["A"].name, got["A"].price, got["A"].specialRule.price, got["A"].specialRule.quantity)
		}
	}
}

func TestCalculateTotalPrice(t *testing.T) {
	for _, c := range []struct {
		want      int
		item      []string
		file_path string
	}{
		{12, []string{"A", "A", "A", "A", "B", "C", "D"}, "tmp/pricing_rules.txt"},
	} {
		got := calculateTotalPrice(c.item, c.file_path)
		if got != c.want {
			t.Errorf("test failed, expected %d, got %d", c.want, got)
		}
	}
}
