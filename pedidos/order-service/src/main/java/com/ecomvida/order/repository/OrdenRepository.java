package com.ecomvida.order.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import com.ecomvida.order.model.Orden;

public interface OrdenRepository extends JpaRepository<Orden, Long> {
}
