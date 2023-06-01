#!/bin/bash
test_path=$(cd "${BASH_SOURCE[0]%/*}"; pwd)
root_path=$(cd "$test_path/../.."; pwd)

. "$root_path/tap.sh" 2>/dev/null || . "$root_path/test/tap.sh"

# Web scraping price from ishares.com
# https://www.ishares.com/uk/individual/en/products/287340/ishares-treasury-bond-1-3yr-ucits-etf
test_ibtausd() {
    curl() {
        echo '<li class="navAmount " data-col="fundHeader.fundNav.navAmount" data-path="">
<span class="header-nav-label navAmount">
NAV as of 09/Jan/2023
</span>
<span class="header-nav-data">
USD 5.43
</span>
<span class="header-info-bubble">
</span>
<br>
<span class="fiftyTwoWeekData">
52 WK: 5.11 - 5.37
</span>
</li>
200'
    }
    export -f curl
    assert "IBTAUSD HTML scrape for price" match "5.43" < <($root_path/bin/setzer price ibtausd)

    curl() {
        echo '<li class="navAmount " data-col="fundHeader.fundNav.navAmount" data-path="">
<span class="header-nav-label navAmount">
NAV as of 09/Jan/2023
</span>
<span>
USD 5.43
</span>
<span class="header-info-bubble">
</span>
<br>
<span class="fiftyTwoWeekData">
52 WK: 5.11 - 5.37
</span>
</li>
200
'
    }
    export -f curl
    assert "IBTAUSD HTML scrape for price not found" fail $root_path/bin/setzer price ibtausd

    curl() {
        echo '<li class="navAmount " data-col="fundHeader.fundNav.navAmount" data-path="">
<span class="header-nav-label navAmount">
NAV as of 09/Jan/2023
</span>
<span class="header-info-bubble">
</span>
<br>
<span class="fiftyTwoWeekData">
52 WK: 5.11 - 5.37
</span>
</li>
200
'
    }
    export -f curl
    assert "IBTAUSD missing price" fail $root_path/bin/setzer price ibtausd

    curl() {
        echo '<li class="navAmount " data-col="fundHeader.fundNav.navAmount" data-path="">
<span class="header-nav-label navAmount">
NAV as of 09/Jan/2023
</span>
<span class="header-nav-data">
USD 12.34567
</span>
<span class="header-info-bubble">
</span>
<br>
<span class="fiftyTwoWeekData">
52 WK: 5.11 - 5.37
</span>
</li>
200
'
    }
    export -f curl
    assert "IBTAUSD price formatting" match "12.34567" < <($root_path/bin/setzer price ibtausd)

    curl() {
       echo '<li class="navAmount " data-col="fundHeader.fundNav.navAmount" data-path="">
<span class="header-nav-label navAmount">
NAV as of 09/Jan/2023
</span>
<span class="header-nav-data">
USD 12.34567
</span>
<span class="header-info-bubble">
</span>
<br>
<span class="fiftyTwoWeekData">
52 WK: 5.11 - 5.37
</span>
</li>
404'
    }
    export -f curl
    assert "IBTAUSD 404 fails" fail $root_path/bin/setzer price ibtausd

    unset -f curl
} && test_ibtausd
