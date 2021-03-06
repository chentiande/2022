create database if not exists ${hiveconf:database};
use ${hiveconf:database};
drop table if exists dm_ct_xdr_lte_5g_s1u_n3_ohvd_area_d;
create EXTERNAL table IF NOT EXISTS dm_ct_xdr_lte_5g_s1u_n3_ohvd_area_d(
`starttime` timestamp,
`province_id` int,
`province_name` varchar(128),
`area_id` int,
`area_name` varchar(255),
`ue_ambr_ul` bigint,
`ue_ambr_dl` bigint,
`apn_ambr_ul` bigint,
`apn_ambr_dl` bigint,
`duration` bigint,
`output_octets` bigint,
`input_octets` bigint,
`total_data` bigint,
`output_packet` bigint,
`input_packet` bigint,
`ul_ipfrag_packs` bigint,
`dl_ipfrag_packs` bigint,
`ul_tcp_retrans_packet` bigint,
`dl_tcp_retrans_packet` bigint,
`ul_tcp_reordering_packet` bigint,
`dl_tcp_reordering_packet` bigint,
`tcp_setup_response_delay` bigint,
`tcp_setup_response_delay_cnt` bigint,
`tcp_setup_ack_delay` bigint,
`tcp_setup_ack_delay_cnt` bigint,
`delay_setup_first_transaction` bigint,
`delay_setup_first_transaction_cnt` bigint,
`delay_first_transaction_first_respackt` bigint,
`delay_first_transaction_first_respackt_cnt` bigint,
`tcp_connect_request_cnt` bigint,
`tcp_connect_succ_cnt` bigint,
`tcp_connect_fail_cnt` bigint,
`tcp_server_response_cnt` bigint,
`tcp_server_response_succ_cnt` bigint,
`tcp_termeral_response_cnt` bigint,
`tcp_termeral_response_succ_cnt` bigint,
`http_bigflow_duration` bigint,
`http_bigflow_duration_cnt` bigint,
`http_bigflow_input_octets` bigint,
`http_smalflow_duration` bigint,
`http_smalflow_duration_cnt` bigint,
`http_smalflow_input_octets` bigint,
`http_download_low_cnt` bigint,
`http_dl_rate_data` bigint,
`http_dl_rate_data_cnt` bigint,
`http_ul_rate_data` bigint,
`http_ul_rate_data_cnt` bigint,
`tcp_connect_delay` bigint,
`tcp_connect_delay_cnt` bigint,
`tcp_termeral_delay_cnt` bigint,
`tcp_server_delay_cnt` bigint,
`game_conn_delay` bigint,
`game_conn_delay_cnt` bigint,
`http_tcp_server_delay` bigint,
`http_tcp_server_delay_cnt` bigint,
`http_tcp_termeral_delay` bigint,
`http_tcp_termeral_delay_cnt` bigint,
`video_tcp_server_delay` bigint,
`video_tcp_server_delay_cnt` bigint,
`video_tcp_termeral_delay` bigint,
`video_tcp_termeral_delay_cnt` bigint,
`im_tcp_tcp_server_delay` bigint,
`im_tcp_tcp_server_delay_cnt` bigint,
`im_tcp_termeral_delay` bigint,
`im_tcp_termeral_delay_cnt` bigint,
`game_tcp_server_delay` bigint,
`game_tcp_server_delay_cnt` bigint,
`game_tcp_termeral_delay` bigint,
`game_tcp_termeral_delay_cnt` bigint,
`chkhand_good` bigint,
`chkhand_bad` bigint,
`rtt_ul_server_good` bigint,
`rtt_ul_server_bat` bigint,
`rtt_dl_termeral_good` bigint,
`rtt_dl_termeral_bat` bigint,
`ul_packet_loss_rate_good_cnt` bigint,
`ul_packet_loss_rate_bat_cnt` bigint,
`dl_packet_loss_rate_good_cnt` bigint,
`dl_packet_loss_rate_bat_cnt` bigint,
`n3_other_cnt` bigint,
`ul_upload_good` bigint,
`ul_upload_bat` bigint,
`dl_download_good` bigint,
`dl_download_bat` bigint,
`response_delay` bigint,
`response_delay_cnt` bigint,
`http_request_cnt` bigint,
`http_succ_cnt` bigint,
`http_noresponse_cnt` bigint,
`http_response_fail_cnt` bigint,
`http_server_fail_cnt` bigint,
`http_termeral_fail_cnt` bigint,
`page_open_succ_cnt` bigint,
`http_respdelay` bigint,
`http_respdelay_cnt` bigint,
`tcp_request_cnt` bigint,
`tcp_sr_request_succ_cnt` bigint,
`tcp_sr_delay` bigint,
`tcp_sr_delay_cnt` bigint,
`tcp_send_delay` bigint,
`tcp_send_delay_cnt` bigint,
`tcp_receive_delay` bigint,
`tcp_receive_delay_cnt` bigint,
`datapack_duration` bigint,
`datapack_duration_cnt` bigint,
`http_bigflow_delay` bigint,
`http_bigflow_delay_cnt` bigint,
`http_smalflow_delay` bigint,
`http_smalflow_delay_cnt` bigint,
`tcp_sr_request_cnt` bigint,
`f_packets_response_delay` bigint,
`f_packets_response_delay_cnt` bigint,
`dns_reply_delay` bigint,
`dns_reply_delay_cnt` bigint,
`first_tcp_pageopen_delay_nodns` bigint,
`first_tcp_pageopen_delay_nodns_cnt` bigint,
`first_tcp_pageopen_delay` bigint,
`first_tcp_pageopen_delay_cnt` bigint,
`first_screen_delay_nodns` bigint,
`first_screen_delay_nodns_cnt` bigint,
`first_screen_delay` bigint,
`first_screen_delay_cnt` bigint,
`first_screen_delaygood` bigint,
`first_screen_delaybad` bigint,
`f_packets_response_delaygood` bigint,
`f_packets_response_delaybad` bigint,
`pageopen_delaygood` bigint,
`pageopen_delaybad` bigint,
`http_dns_reply_delaygood` bigint,
`http_dns_reply_delaybad` bigint,
`connect_good` bigint,
`connect_bad` bigint,
`dns_reply_delaygood` bigint,
`dns_reply_delaybad` bigint,
`n3_http_cnt` bigint,
`videoplay_wait_duration` bigint,
`videoplay_wait_duration_cnt` bigint,
`videoplay_request_cnt` bigint,
`videoplay_succ_cnt` bigint,
`videoplay_halt_cnt` bigint,
`video_recover_duration` bigint,
`video_recover_duration_cnt` bigint,
`video_size` bigint,
`video_dl_duration` bigint,
`video_dl_duration_cnt` bigint,
`video_throughput` bigint,
`video_throughput_cnt` bigint,
`video_bit_rate` bigint,
`video_bit_rate_cnt` bigint,
`video_cache_throughput` bigint,
`video_cache_throughput_cnt` bigint,
`video_halt_rate` bigint,
`video_haltr_ate_cnt` bigint,
`video_avgspeed_good` bigint,
`video_avgspeed_bad` bigint,
`videoplay_halt_good` bigint,
`videoplay_halt_bad` bigint,
`videobit_code_good` bigint,
`videobit_code_bad` bigint,
`first_data_pkg_time` bigint,
`first_data_pkg_time_cnt` bigint,
`last_data_pkg_time` bigint,
`last_data_pkg_time_cnt` bigint,
`cached_video_duration` bigint,
`cached_video_duration_cnt` bigint,
`n3_video_cnt` bigint,
`tcp_syn_cnt` bigint,
`request_cnt` bigint,
`response_cnt` bigint,
`auth_cnt` bigint,
`additional_cnt` bigint,
`dns_resp_succ_cnt` bigint,
`dns_resp_fail_cnt` bigint,
`dns_resp_fail_1_cnt` bigint,
`dns_resp_fail_2_cnt` bigint,
`dns_resp_fail_3_cnt` bigint,
`dns_resp_fail_4_cnt` bigint,
`dns_resp_fail_5_cnt` bigint,
`dns_resp_fail_6_cnt` bigint,
`dns_resp_fail_255_cnt` bigint,
`dns_resp_duration` bigint,
`dns_resp_duration_cnt` bigint,
`dns_delay_less_99ms_cnt` bigint,
`dns_delay_100ms_199ms_cnt` bigint,
`dns_delay_200ms_299ms_cnt` bigint,
`dns_delay_300ms_399ms_cnt` bigint,
`dns_delay_400ms_499ms_cnt` bigint,
`dns_delay_500ms_599ms_cnt` bigint,
`dns_delay_600ms_699ms_cnt` bigint,
`dns_delay_700ms_1s_cnt` bigint,
`dns_delay_1s_2s_cnt` bigint,
`dns_resp_more_2s_cnt` bigint,
`n3_dns_cnt` bigint,
`n3_cnt` bigint,
`imsi_cnt` bigint,
`ul_tcp_renum_rat` decimal(12,6),
`dl_tcp_renum_rat` decimal(12,6),
`ul_tcp_discordnum_rat` decimal(12,6),
`dl_tcp_discordnum_rat` decimal(12,6),
`ul_throughput_rat` decimal(12,6),
`dl_throughput_rat` decimal(12,6),
`http_service_success_rat` decimal(12,6),
`page_open_success_rat` decimal(12,6),
`tcp_build_success_rat` decimal(12,6),
`tcp_server_success_rat` decimal(12,6),
`tcp_termeral_success_rat` decimal(12,6),
`udp_packetloss_rat` decimal(12,6),
`http_noresponse_req` decimal(12,6),
`big_http_download_rat` decimal(20,4),
`sm_http_download_rat` decimal(20,4),
`f_packets_response_good_rat` decimal(12,6),
`http_dns_res_good_rat` decimal(12,6),
`page_open_good_rat` decimal(12,6),
`page_screen_good_rat` decimal(12,6),
`im_good_rat` decimal(12,6),
`game_good_rat` decimal(12,6),
`vide_speed_good_rat` decimal(12,6),
`vide_halt_good_rat` decimal(12,6),
`bit_code_rate_good_rat` decimal(12,6),
`http_service_good_rat` decimal(12,6),
`vide_service_good_rat` decimal(12,6),
`rtt_ul_server_good_rat` decimal(12,6),
`rtt_dl_termeral_good_rat` decimal(12,6),
`rtt_service_good_rat` decimal(12,6),
`ul_upload_good_rat` decimal(12,6),
`dl_download_good_rat` decimal(12,6),
`tran_service_good_rate` decimal(12,6),
`udp_packetaccept_good_rat` decimal(12,6),
`service_perceive_good_rat` decimal(12,6),
`dns_resp_succ_rate` decimal(12,6)) 
PARTITIONED BY ( 
  `p_date` string)
ROW FORMAT DELIMITED 
  FIELDS TERMINATED BY '|' NULL DEFINED AS '' 
location "${hivevar:basepath}/dm_ct_xdr_lte_5g_s1u_n3_ohvd_area_d/";