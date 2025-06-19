package com.ecomvida.address.security;

import io.jsonwebtoken.*;
import jakarta.servlet.*;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;

@Component
public class JwtFilter extends OncePerRequestFilter {

    @Value("${jwt.secret}")
    private String secret;

    private static final String AUTH_HEADER = "Authorization";
    public static final String ATTR_USUARIO_ID = "usuarioId";

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain)
            throws ServletException, IOException {

        String header = request.getHeader(AUTH_HEADER);
        if (header != null && header.startsWith("Bearer ")) {
            String token = header.substring(7);
            try {
                Claims claims = Jwts.parser()
                        .setSigningKey(secret)
                        .parseClaimsJws(token)
                        .getBody();

                request.setAttribute(ATTR_USUARIO_ID, claims.get("id"));
            } catch (JwtException e) {
                throw new ServletException("Token inválido");
            }
        } else {
            throw new ServletException("Falta el token en Authorization header");
        }

        chain.doFilter(request, response);
    }
}
