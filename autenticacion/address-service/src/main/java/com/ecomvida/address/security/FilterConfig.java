package com.ecomvida.address.security;

import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class FilterConfig {

    @Bean
public FilterRegistrationBean<JwtFilter> jwtFilterBean() {
    FilterRegistrationBean<JwtFilter> registrationBean = new FilterRegistrationBean<>();
    registrationBean.setFilter(new JwtFilter());
    registrationBean.addUrlPatterns("/api/*");
    return registrationBean;
}


    
}
