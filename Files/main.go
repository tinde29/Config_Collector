package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	
	"github.com/PuerkitoBio/goquery"
)

var client = &http.Client{}

func main() {
	channels := []string{
    "https://t.me/s/Airdorap_Free",
    "https://t.me/s/Awlix_V2RAY",
    "https://t.me/s/CUSTOMVPNSERVER",
    "https://t.me/s/Capoit",
    "https://t.me/s/ConfigsHubPlus",
    "https://t.me/s/Configforvpn01",
    "https://t.me/s/DailyV2RY",
    "https://t.me/s/DarkTeam_VPN",
    "https://t.me/s/DarkVPNpro",
    "https://t.me/s/DeamNet_Proxy",
    "https://t.me/s/EUT_VPN",
    "https://t.me/s/EliV2ray",
    "https://t.me/s/FalconPolV2rayNG",
    "https://t.me/s/FoXrayIran",
    "https://t.me/s/FreeNet1500",
    "https://t.me/s/Free_Internet_Iran",
    "https://t.me/s/Good_V2rayy",
    "https://t.me/s/GozargahVPN",
    "https://t.me/s/Helix_Servers",
    "https://t.me/s/IBv2ray",
    "https://t.me/s/IRN_VPN",
    "https://t.me/s/LuminousNet",
    "https://t.me/s/MrV2Ray",
    "https://t.me/s/OpenSSTPVpn",
    "https://t.me/s/Outline_ir",
    "https://t.me/s/PAINB0Y",
    "https://t.me/s/PASARGAD_V2rayNG",
    "https://t.me/s/PROXY_MTM",
    "https://t.me/s/ProxyForOpeta",
    "https://t.me/s/Qv2rayDONATED",
    "https://t.me/s/ROMAX_VPN",
    "https://t.me/s/SINABIGO",
    "https://t.me/s/SiNABiGO",
    "https://t.me/s/Speeds_vpn1",
    "https://t.me/s/SvnTeam",
    "https://t.me/s/TIKTOK_PROXY",
    "https://t.me/s/TUICity",
    "https://t.me/s/TVCminer",
    "https://t.me/s/UnlimitedDev",
    "https://t.me/s/V2FETCH",
    "https://t.me/s/V2RAY_NEW",
    "https://t.me/s/V2Ray_FreedomIran",
    "https://t.me/s/V2RayOxygen",
    "https://t.me/s/V2rayCollector",
    "https://t.me/s/V2rayCollectorDonate",
    "https://t.me/s/V2rayNG_im",
    "https://t.me/s/V2rayNGvp",
    "https://t.me/s/V2rayNGvpni",
    "https://t.me/s/V2raysFree",
    "https://t.me/s/VPNGate_Config",
    "https://t.me/s/VPN_443",
    "https://t.me/s/VpnProSec",
    "https://t.me/s/VpnSkyy",
    "https://t.me/s/VpnWLF",
    "https://t.me/s/Vpn_Mikey",
    "https://t.me/s/Vpn_Xpace",
    "https://t.me/s/WPSNET",
    "https://t.me/s/Watashi_VPN",
    "https://t.me/s/WeePeeN",
    "https://t.me/s/WomanLifeFreedomVPN",
    "https://t.me/s/XpnTeam",
    "https://t.me/s/alfred_config",
    "https://t.me/s/allv2ray",
    "https://t.me/s/alo_v2rayng",
    "https://t.me/s/angus_vpn",
    "https://t.me/s/antifilterservice",
    "https://t.me/s/appsooner",
    "https://t.me/s/arv2ray",
    "https://t.me/s/asak_vpn",
    "https://t.me/s/asintech",
    "https://t.me/s/aspeedvpn",
    "https://t.me/s/astrovpn_official",
    "https://t.me/s/awlix_ir",
    "https://t.me/s/azadnet",
    "https://t.me/s/azarbayjab1",
    "https://t.me/s/beiten",
    "https://t.me/s/bermudavpn24",
    "https://t.me/s/bglvps",
    "https://t.me/s/bigsmoke_config",
    "https://t.me/s/blueberrynetwork",
    "https://t.me/s/bored_vpn",
    "https://t.me/s/bypass_filter",
    "https://t.me/s/catvpns",
    "https://t.me/s/cconfig_v2ray",
    "https://t.me/s/channels_assets",
    "https://t.me/s/city_v2rayng",
    "https://t.me/s/client_proo",
    "https://t.me/s/club_profsor",
    "https://t.me/s/config_v2ray",
    "https://t.me/s/configforvpn",
    "https://t.me/s/configpositive",
    "https://t.me/s/configt",
    "https://t.me/s/configv2rayforfree",
    "https://t.me/s/custom_config",
    "https://t.me/s/customizev2ray",
    "https://t.me/s/cvrnet",
    "https://t.me/s/dailyv2ry",
    "https://t.me/s/daredevill_404",
    "https://t.me/s/deragv2ray",
    "https://t.me/s/digiv2ray",
    "https://t.me/s/donald_vpn",
    "https://t.me/s/drvpn_net",
    "https://t.me/s/easy_free_vpn",
    "https://t.me/s/eliya_chiter0",
    "https://t.me/s/entrynet",
    "https://t.me/s/ev2rayy",
    "https://t.me/s/expressvpn_420",
    "https://t.me/s/external_net",
    "https://t.me/s/farahvpn",
    "https://t.me/s/fasst_vpn",
    "https://t.me/s/fast_2ray",
    "https://t.me/s/fastkanfig",
    "https://t.me/s/fastshadow_vpn",
    "https://t.me/s/filterk0sh",
    "https://t.me/s/flyv2ray",
    "https://t.me/s/forwardv2ray",
    "https://t.me/s/foxrayiran",
    "https://t.me/s/freakconfig",
    "https://t.me/s/freakconfig1",
    "https://t.me/s/freakconfig2",
    "https://t.me/s/free4allVPN",
    "https://t.me/s/freeconfigv2",
    "https://t.me/s/freeconfigvpns",
    "https://t.me/s/freeiranweb",
    "https://t.me/s/freenapsternetv",
    "https://t.me/s/freev2raym",
    "https://t.me/s/freev2rayssr",
    "https://t.me/s/freevirgoolnet",
    "https://t.me/s/frev2rayng",
    "https://t.me/s/fsv2ray",
    "https://t.me/s/ghalagyann",
    "https://t.me/s/godv2ray_ng",
    "https://t.me/s/godot404",
    "https://t.me/s/golestan_vpn",
    "https://t.me/s/grizzlyvpn",
    "https://t.me/s/hajimamadvpn",
    "https://t.me/s/hamster_vpnn",
    "https://t.me/s/hatunnel_vpn",
    "https://t.me/s/hcv2ray",
    "https://t.me/s/hope_net",
    "https://t.me/s/hopev2ray",
    "https://t.me/s/hormozvpn",
    "https://t.me/s/hose_io",
    "https://t.me/s/https_config_injector",
    "https://t.me/s/iSegaro",
    "https://t.me/s/icv2ray",
    "https://t.me/s/idigitalz",
    "https://t.me/s/imrv2ray",
    "https://t.me/s/imtproxy_ir",
    "https://t.me/s/inikotesla",
    "https://t.me/s/ios_v2",
    "https://t.me/s/ipV2Ray",
    "https://t.me/s/ipcloudflaree",
    "https://t.me/s/ipcloudflaretamiz",
    "https://t.me/s/ipv2ray",
    "https://t.me/s/irancpi_vpn",
    "https://t.me/s/iranbaxvpn",
    "https://t.me/s/iraniv2ray_config",
    "https://t.me/s/irv2rey",
    "https://t.me/s/isvvpn",
    "https://t.me/s/kafing_2",
    "https://t.me/s/kiava",
    "https://t.me/s/kingspeedchanel",
    "https://t.me/s/kingofilter",
    "https://t.me/s/liq_VPN",
    "https://t.me/s/ln2ray",
    "https://t.me/s/lombo_channel",
    "https://t.me/s/mahdiserver",
    "https://t.me/s/mahsaamoon1",
    "https://t.me/s/mahvarehnewssat",
    "https://t.me/s/manzariyeh_rasht",
    "https://t.me/s/marambashi",
    "https://t.me/s/maznet",
    "https://t.me/s/mehrosaboran",
    "https://t.me/s/meli_proxyy",
    "https://t.me/s/mester_v2ray",
    "https://t.me/s/mftizi",
    "https://t.me/s/mgvpnsale",
    "https://t.me/s/mikasavpn",
    "https://t.me/s/miov2ray",
    "https://t.me/s/moftinet",
    "https://t.me/s/molovpn",
    "https://t.me/s/msv2raynp",
    "https://t.me/s/n2vpn",
    "https://t.me/s/napsternetv_config",
    "https://t.me/s/netmellianti",
    "https://t.me/s/networknim",
    "https://t.me/s/new_proxy_channel",
    "https://t.me/s/nofiltering2",
    "https://t.me/s/noforcedheaven",
    "https://t.me/s/npvv2rayfilter",
    "https://t.me/s/nufilter",
    "https://t.me/s/ohvpn",
    "https://t.me/s/orange_vpns",
    "https://t.me/s/outline_vpn",
    "https://t.me/s/outlinevpnir",
    "https://t.me/s/pars_vpn3",
    "https://t.me/s/pashmam_vpn",
    "https://t.me/s/piazshekan",
    "https://t.me/s/pishiserver",
    "https://t.me/s/privatevpns",
    "https://t.me/s/proprojec",
    "https://t.me/s/proxiiraniii",
    "https://t.me/s/proxy_kafee",
    "https://t.me/s/proxy_mtm",
    "https://t.me/s/proxy_n1",
    "https://t.me/s/proxyfull",
    "https://t.me/s/prroxyng",
    "https://t.me/s/puni_shop_v2rayng",
    "https://t.me/s/qeshmserver",
    "https://t.me/s/raze_vpn",
    "https://t.me/s/realvpnmaster",
    "https://t.me/s/rnrifci",
    "https://t.me/s/satoshivpn",
    "https://t.me/s/savagev2ray",
    "https://t.me/s/selinc",
    "https://t.me/s/serveriran11",
    "https://t.me/s/servermomo",
    "https://t.me/s/shadoowvpnn",
    "https://t.me/s/shadowsocksshop",
    "https://t.me/s/shokhmiplus",
    "https://t.me/s/sinabigo",
    "https://t.me/s/sinavm",
    "https://t.me/s/sobi_vpn",
    "https://t.me/s/special_net8",
    "https://t.me/s/spikevpn",
    "https://t.me/s/srcvpn",
    "https://t.me/s/summertimeus",
    "https://t.me/s/superv2rang",
    "https://t.me/s/svnteam",
    "https://t.me/s/talentvpn",
    "https://t.me/s/tehranargo",
    "https://t.me/s/tehranargo1",
    "https://t.me/s/thexconfig",
    "https://t.me/s/thunderv2ray",
    "https://t.me/s/tv_v2ray",
    "https://t.me/s/v2ary",
    "https://t.me/s/v2aryng_vpn",
    "https://t.me/s/v2boxvpnn",
    "https://t.me/s/v2fre",
    "https://t.me/s/v2graphy",
    "https://t.me/s/v2line",
    "https://t.me/s/v2logy",
    "https://t.me/s/v2net_iran",
    "https://t.me/s/v2ngfast",
    "https://t.me/s/v2ra2",
    "https://t.me/s/v2raand",
    "https://t.me/s/v2rang00",
    "https://t.me/s/v2range",
    "https://t.me/s/v2raxx",
    "https://t.me/s/v2ray6388",
    "https://t.me/s/v2rayNGNeT",
    "https://t.me/s/v2rayNG_Matsuri",
    "https://t.me/s/v2rayNG_VPNN",
    "https://t.me/s/v2ray_alpha07",
    "https://t.me/s/v2ray_ar",
    "https://t.me/s/v2ray_configs_pool",
    "https://t.me/s/v2ray_fark",
    "https://t.me/s/v2ray_for_free",
    "https://t.me/s/v2ray_ng",
    "https://t.me/s/v2ray_one1",
    "https://t.me/s/v2ray_raha",
    "https://t.me/s/v2ray_rolly",
    "https://t.me/s/v2rayan",
    "https://t.me/s/v2rayargon",
    "https://t.me/s/v2raych",
    "https://t.me/s/v2rayfast",
    "https://t.me/s/v2rayfast_7",
    "https://t.me/s/v2rayfree_irr",
    "https://t.me/s/v2rayiman",
    "https://t.me/s/v2raylandd",
    "https://t.me/s/v2rayn2g",
    "https://t.me/s/v2rayn_server",
    "https://t.me/s/v2rayng3",
    "https://t.me/s/v2rayng_city",
    "https://t.me/s/v2rayng_fa2",
    "https://t.me/s/v2rayng_madam",
    "https://t.me/s/v2rayng_prime",
    "https://t.me/s/v2rayng_v",
    "https://t.me/s/v2rayngvpn",
    "https://t.me/s/v2rayngvpnn",
    "https://t.me/s/v2rayngzendegimamad",
    "https://t.me/s/v2rayprotocol",
    "https://t.me/s/v2rayyngvpn",
    "https://t.me/s/v2rez",
    "https://t.me/s/v2rray_ng",
    "https://t.me/s/v2ry_proxy",
    "https://t.me/s/v2ryng01",
    "https://t.me/s/v2ryng_vpn",
    "https://t.me/s/v2ryngfree",
    "https://t.me/s/v2safe",
    "https://t.me/s/v2safee",
    "https://t.me/s/v_2rayngvpn",
    "https://t.me/s/vipvpn_v2ray",
    "https://t.me/s/vistav2ray",
    "https://t.me/s/vmesc",
    "https://t.me/s/vmess_vless_pro",
    "https://t.me/s/vmessiran",
    "https://t.me/s/vmesskhodam",
    "https://t.me/s/vmesskhodam_vip",
    "https://t.me/s/vmessorg",
    "https://t.me/s/vmessprotocol",
    "https://t.me/s/vmessq",
    "https://t.me/s/vp22ray",
    "https://t.me/s/vpnaloo",
    "https://t.me/s/vpn_go67",
    "https://t.me/s/vpn_ioss",
    "https://t.me/s/vpnazadland",
    "https://t.me/s/vpnconfignet",
    "https://t.me/s/vpnfail_v2ray",
    "https://t.me/s/vpnhubmarket",
    "https://t.me/s/vpnkanfik",
    "https://t.me/s/vpnowl",
    "https://t.me/s/vpnsaint",
    "https://t.me/s/vpnstorefast",
    "https://t.me/s/vpnv2rayngv",
    "https://t.me/s/vpnxyam_ir",
    "https://t.me/s/wedbaztel",
    "https://t.me/s/wsbvpn",
    "https://t.me/s/xpnteam",
    "https://t.me/s/xrayproxy",
    "https://t.me/s/xvproxy",
    "https://t.me/s/zen_cloud",
    "https://t.me/s/zede_filteri",
    "https://t.me/s/zibanabz",
    "https://t.me/s/zohalserver",
    "https://t.me/v2rayNGcloud",
    "https://t.me/s/AM_TEAMMM",
    "https://t.me/s/ARv2ray",
    "https://t.me/s/Awlix_ir",
    "https://t.me/s/BestV2rang",
    "https://t.me/s/Capital_NET",
    "https://t.me/s/CloudCityy",
    "https://t.me/s/ConfigsHUB",
    "https://t.me/s/Cov2ray",
    "https://t.me/s/DigiV2ray",
    "https://t.me/s/DirectVPN",
    "https://t.me/s/Easy_Free_VPN",
    "https://t.me/s/FOX_VPN66",
    "https://t.me/s/FreakConfig",
    "https://t.me/s/FreeV2rays",
    "https://t.me/s/FreeVlessVpn",
    "https://t.me/s/God_CONFIG",
    "https://t.me/s/HTTPCustomLand",
    "https://t.me/s/Hope_Net",
    "https://t.me/s/INIT1984",
    "https://t.me/s/L_AGVPN13",
    "https://t.me/s/LoRd_uL4mo",
    "https://t.me/s/Lockey_vpn",
    "https://t.me/s/MTConfig",
    "https://t.me/s/MehradLearn",
    "https://t.me/s/MsV2ray",
    "https://t.me/s/NIM_VPN_ir",
    "https://t.me/s/NetAccount",
    "https://t.me/s/Network_442",
    "https://t.me/s/OutlineVpnOfficial",
    "https://t.me/s/Outline_Vpn",
    "https://t.me/s/Outlinev2rayNG",
    "https://t.me/s/Parsashonam",
    "https://t.me/s/PrivateVPNs",
    "https://t.me/s/Proxy_PJ",
    "https://t.me/s/Qv2raychannel",
    "https://t.me/s/Royalping_ir",
    "https://t.me/s/SafeNet_Server",
    "https://t.me/s/ServerNett",
    "https://t.me/s/ShadowSocks_s",
    "https://t.me/s/ShadowsocksM",
    "https://t.me/s/TPvpn_Official",
    "https://t.me/s/V2RayChannel",
    "https://t.me/s/V2RayTz",
    "https://t.me/s/V2parsin",
    "https://t.me/s/V2pedia",
    "https://t.me/s/V2rayNGmat",
    "https://t.me/s/V2rayNGn",
    "https://t.me/s/V2rayng_Fast",
    "https://t.me/s/V2rayngninja",
    "https://t.me/s/VPNCLOP",
    "https://t.me/s/VPNCUSTOMIZE",
    "https://t.me/s/VpnFreeSec",
    "https://t.me/s/VmessProtocol",
    "https://t.me/s/VorTexIRN",
    "https://t.me/s/WebShecan",
    "https://t.me/s/XsV2ray",
    "https://t.me/s/YtTe3la",
    "https://t.me/s/azadi_az_inja_migzare",
    "https://t.me/s/configV2rayForFree",
    "https://t.me/s/configV2rayNG",
    "https://t.me/s/custom_14",
    "https://t.me/s/customv2ray",
    "https://t.me/s/daorzadannet",
    "https://t.me/s/fnet00",
    "https://t.me/s/free1_vpn",
    "https://t.me/s/free_v2rayyy",
    "https://t.me/s/freeland8",
    "https://t.me/s/frev2ray",
    "https://t.me/s/hashmakvpn",
    "https://t.me/s/iP_CF",
    "https://t.me/s/iTV2RAY",
    "https://t.me/s/iran_ray",
    "https://t.me/s/iranvpnet",
    "https://t.me/s/lightning6",
    "https://t.me/s/lrnbymaa",
    "https://t.me/s/melov2ray",
    "https://t.me/s/oneclickvpnkeys",
    "https://t.me/s/ovpn2",
    "https://t.me/s/polproxy",
    "https://t.me/s/prrofile_purple",
    "https://t.me/s/proxyymeliii",
    "https://t.me/s/rayvps",
    "https://t.me/s/reality_daily",
    "https://t.me/s/rxv2ray",
    "https://t.me/s/shh_proxy",
    "https://t.me/s/shopingv2ray",
    "https://t.me/s/therealaleph",
    "https://t.me/s/ultrasurf_12",
    "https://t.me/s/v2Line",
    "https://t.me/s/v2_team",
    "https://t.me/s/v2_vmess",
    "https://t.me/s/v2ray1_ng",
    "https://t.me/s/v2rayNG_VPN",
    "https://t.me/s/v2ray_swhil",
    "https://t.me/s/v2ray_vpn_ir",
    "https://t.me/s/v2rayan",
    "https://t.me/s/v2rayng_config_amin",
    "https://t.me/s/v2rayng_org",
    "https://t.me/s/v2rayng_vpnrog",
    "https://t.me/s/v2ray_outlineir",
    "https://t.me/s/vless_vmess",
    "https://t.me/s/vlessconfig",
    "https://t.me/s/vmess_vless_v2rayng",
    "https://t.me/s/vpn_ocean",
    "https://t.me/s/vpn_proxy_custom",
    "https://t.me/s/vpn_tehran",
    "https://t.me/s/vpn_xw",
    "https://t.me/s/vpnmasi",
    "https://t.me/s/vip_vpn_2022",
    "https://t.me/s/yaney_01",
}

	configs := map[string]string{
		"ss":     "",
		"vmess":  "",
		"trojan": "",
		"vless":  "",
		"mixed":  "",
	}

	myregex := map[string]string{
		"ss":     `(.{3})ss:\/\/`,
		"vmess":  `vmess:\/\/`,
		"trojan": `trojan:\/\/`,
		"vless":  `vless:\/\/`,
	}

	for i := 0; i < len(channels); i++ {
		all_messages := false
		if strings.Contains(channels[i], "{all_messages}") {
			all_messages = true
			channels[i] = strings.Split(channels[i], "{all_messages}")[0]
		}

		req, err := http.NewRequest("GET", channels[i], nil)
		if err != nil {
			log.Fatalf("Error When requesting to: %s Error : %s", channels[i], err)
		}

		resp, err1 := client.Do(req)
		if err1 != nil {
			log.Fatal(err1)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		messages := doc.Find(".tgme_widget_message_wrap").Length()
		link, exist := doc.Find(".tgme_widget_message_wrap .js-widget_message").Last().Attr("data-post")
		if messages < 100 && exist == true {
			number := strings.Split(link, "/")[1]
			fmt.Println(number)

			doc = GetMessages(10, doc, number, channels[i])
		}

		if all_messages {
			fmt.Println(doc.Find(".js-widget_message_wrap").Length())
			doc.Find(".tgme_widget_message_text").Each(func(j int, s *goquery.Selection) {
				message_text := s.Text()
				lines := strings.Split(message_text, "\n")
				for a := 0; a < len(lines); a++ {
					for _, regex_value := range myregex {
						re := regexp.MustCompile(regex_value)
						lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
							return "\n" + match
						})
					}
					for proto := range configs {
						if strings.Contains(lines[a], proto) {
							configs["mixed"] += "\n" + lines[a] + "\n"
						}
					}
				}
			})
		} else {
			doc.Find("code,pre").Each(func(j int, s *goquery.Selection) {
				message_text := s.Text()
				lines := strings.Split(message_text, "\n")
				for a := 0; a < len(lines); a++ {
					for proto_regex, regex_value := range myregex {
						re := regexp.MustCompile(regex_value)
						lines[a] = re.ReplaceAllStringFunc(lines[a], func(match string) string {
							if proto_regex == "ss" {
								if match[:3] == "vme" {
									return "\n" + match
								} else if match[:3] == "vle" {
									return "\n" + match
								} else {
									return "\n" + match
								}
							} else {
								return "\n" + match
							}
						})

						if len(strings.Split(lines[a], "\n")) > 1 {
							myconfigs := strings.Split(lines[a], "\n")
							for i := 0; i < len(myconfigs); i++ {
								if myconfigs[i] != "" {
									re := regexp.MustCompile(regex_value)
									myconfigs[i] = strings.ReplaceAll(myconfigs[i], " ", "")
									match := re.FindStringSubmatch(myconfigs[i])
									if len(match) >= 1 {
										if proto_regex == "ss" {
											if match[1][:3] == "vme" {
												configs["vmess"] += "\n" + myconfigs[i] + "\n"
											} else if match[1][:3] == "vle" {
												configs["vless"] += "\n" + myconfigs[i] + "\n"
											} else {
												configs["ss"] += "\n" + myconfigs[i][3:] + "\n"
											}
										} else {
											configs[proto_regex] += "\n" + myconfigs[i] + "\n"
										}
									}
								}
							}
						}
					}
				}
			})
		}
	}

	for proto, configcontent := range configs {
		// 		reverse mode :
		// 		lines := strings.Split(configcontent, "\n")
		// 		reversed := reverse(lines)
		// 		WriteToFile(strings.Join(reversed, "\n"), proto+"_iran.txt")
		// 		simple mode :
		WriteToFile(RemoveDuplicate(configcontent), proto+"_iran.txt")
	}

}

func WriteToFile(fileContent string, filePath string) {

	// Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		// If the file exists, clear its content
		err = ioutil.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			fmt.Println("Error clearing file:", err)
			return
		}
	} else if os.IsNotExist(err) {
		// If the file does not exist, create it
		_, err = os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	} else {
		// If there was some other error, print it and return
		fmt.Println("Error checking file:", err)
		return
	}

	// Write the new content to the file
	err := ioutil.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File written successfully")
}

func load_more(link string) *goquery.Document {
	req, _ := http.NewRequest("GET", link, nil)
	fmt.Println(link)
	resp, _ := client.Do(req)
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	return doc
}

func GetMessages(length int, doc *goquery.Document, number string, channel string) *goquery.Document {
	x := load_more(channel + "?before=" + number)
	html2, _ := x.Html()
	reader2 := strings.NewReader(html2)
	doc2, _ := goquery.NewDocumentFromReader(reader2)
	doc.Find("body").AppendSelection(doc2.Find("body").Children())
	newDoc := goquery.NewDocumentFromNode(doc.Selection.Nodes[0])
	messages := newDoc.Find(".js-widget_message_wrap").Length()
	if messages > length {
		return newDoc
	} else {
		num, _ := strconv.Atoi(number)
		n := num - 21
		if n > 0 {
			ns := strconv.Itoa(n)
			GetMessages(length, newDoc, ns, channel)
		} else {
			return newDoc
		}
	}
	return newDoc
}

func RemoveDuplicate(config string) string {
	lines := strings.Split(config, "\n")

	// Use a map to keep track of unique lines
	uniqueLines := make(map[string]bool)

	// Loop over lines and add unique lines to map
	for _, line := range lines {
		if len(line) > 0 {
			uniqueLines[line] = true
		}
	}

	// Join unique lines into a string
	uniqueString := strings.Join(getKeys(uniqueLines), "\n")

	return uniqueString
}

func getKeys(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func getTimestamp(message string) int64 {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(message), &data); err != nil {
		// Handle the error if necessary
		return 0
	}
	timestamp, _ := data["date"].(int64)
	return timestamp
}
