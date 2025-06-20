package com.ecomvida.brand.ws;

import com.ecomvida.brand.service.MarcaService;
import org.apache.cxf.Bus;
import org.apache.cxf.jaxws.EndpointImpl;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import jakarta.xml.ws.Endpoint;

@Configuration
public class CxfConfig {

    private final MarcaService marcaService;

    public CxfConfig(MarcaService marcaService) {
        this.marcaService = marcaService;
    }

    @Bean
    public MarcaEndpoint marcaEndpoint() {
        return new MarcaEndpoint(marcaService);
    }

    @Bean
    public Endpoint endpoint(Bus bus) {
        EndpointImpl endpoint = new EndpointImpl(bus, marcaEndpoint());
        endpoint.publish("/MarcaEndpoint");
        return endpoint;
    }
}
